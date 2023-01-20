package config

import "time"

type JwtConfig struct {
	AcSign string `mapstructure:"acSign"`
	ReSign string `mapstructure:"reSign"`
	AcExp  int64  `mapstructure:"acExp"`
	AcAlo  int64  `mapstructure:"acAlo"`
	ReExp  int64  `mapstructure:"reExp"`
	Issuer string `mapstructure:"issuer"`
}

// AcExpTime
// @Date 2023-01-20 22:46:41
// @Return time.Duration
// @Method
// @Description: access-token 过期时间
func (j *JwtConfig) AcExpTime() time.Duration {
	return time.Hour * time.Duration(j.AcExp)
}

// AcAllowExpTime
// @Date 2023-01-20 22:51:33
// @Return time.Duration
// @Method
// @Description: 允许过期时间
func (j *JwtConfig) AcAllowExpTime() time.Duration {
	return time.Minute * time.Duration(j.ReExp)
}

// ReExpTime
// @Date 2023-01-20 22:46:55
// @Return time.Duration
// @Method
// @Description: refresh-token 过期时间
func (j *JwtConfig) ReExpTime() time.Duration {
	return time.Hour * time.Duration(j.ReExp)
}
