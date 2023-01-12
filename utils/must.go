package utils

// Must
// @Date 2023-01-12 18:23:11
// @Param f func() error
// @Description: 必须完成某个操作，否则就panic
func Must(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}
