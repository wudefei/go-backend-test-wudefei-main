package util

/**
 * 常用通用的工具处理方法
 *
 **/
func ContainString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
