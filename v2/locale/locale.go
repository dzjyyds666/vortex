package locale

var V = "{\"code_for_internal_error.zh-cn\":\"系统内部故障\",\"code_for_params_invaild.zh-cn\":\"参数错误\",\"code_for_permission_deny.zh-cn\":\"权限不足\",\"code_for_success.en-us\":\"sucess\",\"code_for_success.zh-cn\":\"请求成功\",\"code_for_unauthorized.zh-cn\":\"未授权或token无效\"}"

var K = struct {
	CODE_FOR_SUCCESS string
	CODE_FOR_UNAUTHORIZED string
	CODE_FOR_PARAMS_INVAILD string
	CODE_FOR_PERMISSION_DENY string
	CODE_FOR_INTERNAL_ERROR string
} {
	CODE_FOR_UNAUTHORIZED: "code_for_unauthorized",
	CODE_FOR_PARAMS_INVAILD: "code_for_params_invaild",
	CODE_FOR_PERMISSION_DENY: "code_for_permission_deny",
	CODE_FOR_INTERNAL_ERROR: "code_for_internal_error",
	CODE_FOR_SUCCESS: "code_for_success",
}
