package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	crand "crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	mrand "math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
	"wat.ink/layout/fiber/pkg/config"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

// GenerateSalt 生成指定长度的随机盐值
func GenerateSalt(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mrand.Intn(len(letters))]
	}
	return string(b)
}

// GeneratePassword 生成密码哈希
func GeneratePassword(password, salt string) string {
	// 将密码和盐值拼接
	str := password + salt + config.Conf.MD5.Hash
	
	// 计算 MD5 哈希
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Paginate 计算分页参数
func Paginate(page, pageSize int) (offset, limit int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	
	offset = (page - 1) * pageSize
	limit = pageSize
	return
}

// ValidatePassword 验证密码是否正确
func ValidatePassword(inputPassword, hashedPassword, salt string) bool {
	return GeneratePassword(inputPassword, salt) == hashedPassword
}

// ValidateEmail 验证邮箱格式
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// ValidatePassword 验证密码强度
func ValidatePasswordStrength(password string) bool {
	var (
		hasNumber    = false
		hasLower     = false
		hasUpper     = false
		hasSpecial   = false
		length       = len(password)
	)

	if length < 8 {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasNumber && hasUpper && hasLower && hasSpecial
}

// Sanitize 清理字符串中的特殊字符
func Sanitize(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, s)
}

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mrand.Intn(len(letters))]
	}
	return string(b)
}

// FormatFileSize 格式化文件大小
func FormatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

// TimeFormat 常用时间格式
const (
	TimeFormatDate     = "2006-01-02"
	TimeFormatDateTime = "2006-01-02 15:04:05"
)

// FormatTime 格式化时间
func FormatTime(t time.Time, layout string) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(layout)
}

// ParseTime 解析时间字符串
func ParseTime(timeStr, layout string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}
	return time.Parse(layout, timeStr)
}

// GetStartOfDay 获取一天的开始时间
func GetStartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// GetEndOfDay 获取一天的结束时间
func GetEndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

// EncryptAES AES加密
func EncryptAES(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	
	// 创建随机IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(crand.Reader, iv); err != nil {
		return nil, err
	}
	
	// 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	padded := pkcs7Padding(data, aes.BlockSize)
	encrypted := make([]byte, len(padded))
	mode.CryptBlocks(encrypted, padded)
	
	// 拼接IV和密文
	return append(iv, encrypted...), nil
}

// DecryptAES AES解密
func DecryptAES(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	
	if len(data) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	
	iv := data[:aes.BlockSize]
	encrypted := data[aes.BlockSize:]
	
	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)
	
	return pkcs7UnPadding(decrypted)
}

// pkcs7Padding PKCS7填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding PKCS7去填充
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid padding size")
	}
	padding := int(data[length-1])
	if padding > length {
		return nil, errors.New("invalid padding size")
	}
	return data[:length-padding], nil
} 