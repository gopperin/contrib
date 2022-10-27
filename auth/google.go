package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Google 谷歌动态验证码
type Google struct {
	window int64
}

// New New
func Default() *Google {
	return &Google{1}
}

// New New
func New(window int64) *Google {
	if window < 1 {
		return &Google{1}
	}
	return &Google{window}
}

// un un
func (g *Google) un() int64 {
	return time.Now().UnixNano() / 1000 / 30
}

// hmacSha1 hmacSha1
func (g *Google) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

func (g *Google) base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (g *Google) base32decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func (g *Google) toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func (g *Google) toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func (g *Google) oneTimePassword(key []byte, data []byte) uint32 {
	hash := g.hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := g.toUint32(hashParts)
	return number % 1000000
}

// GetSecret 获取秘钥
func (g *Google) GetSecret() string {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, g.un())
	return strings.ToUpper(g.base32encode(g.hmacSha1(buf.Bytes(), nil)))
}

// SetWindow 设置时间窗
func (g *Google) SetWindow(window int64) {
	g.window = window
}

// GetCode 获取动态码
func (g *Google) GetCode(timestamp int64, secret string) (string, error) {
	secretUpper := strings.ToUpper(secret)
	secretKey, err := g.base32decode(secretUpper)
	if err != nil {
		return "", err
	}
	number := g.oneTimePassword(secretKey, g.toBytes(timestamp/30))
	return fmt.Sprintf("%06d", number), nil
}

// GetQrcode 获取动态码二维码内容
func (g *Google) GetQrcode(user, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
}

// GetQrcodeURL 获取动态码二维码图片地址,这里是第三方二维码api
func (g *Google) GetQrcodeURL(user, secret string) string {
	qrcode := g.GetQrcode(user, secret)
	width := "200"
	height := "200"
	data := url.Values{}
	data.Set("data", qrcode)
	return "https://api.qrserver.com/v1/create-qr-code/?" + data.Encode() + "&size=" + width + "x" + height + "&ecc=M"
}

// VerifyCode 验证动态码
func (g *Google) VerifyCode(secret, code string) (bool, error) {
	_timestamp := time.Now().Unix()

	var i int64
	for i = 0; i <= g.window-1; i++ {

		_t := _timestamp - int64(30*i)
		_code, err := g.GetCode(_t, secret)
		if err != nil {
			continue
		}

		if _code == code {
			return true, nil
		}

	}
	return false, nil
}
