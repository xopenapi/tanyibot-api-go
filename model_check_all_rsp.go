/*
 * Tanyi Bot open api
 *
 * tanyibot api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tanyibot
// CheckAllRsp struct for CheckAllRsp
type CheckAllRsp struct {
	// 响应码
	Code int32 `json:"code,omitempty"`
	// 请求id
	RequestId string `json:"requestId,omitempty"`
	// 响应说明
	ResultMsg string `json:"resultMsg,omitempty"`
	ErrorStackTrace string `json:"errorStackTrace,omitempty"`
	// 任务id
	RobotCallJobId string `json:"robotCallJobId,omitempty"`
	// 任务名
	RobotCallJobName string `json:"robotCallJobName,omitempty"`
	// 创建时间
	CreateTime string `json:"createTime,omitempty"`
	// 任务创建人
	CreateUserName string `json:"createUserName,omitempty"`
	// 任务状态，NOT_STARTED(0, \"未开始\"),IN_PROCESS(1, \"进行中\")，COMPLETED(2, \"已完成\"),RUNNABLE(3, \"可运行\"),USER_PAUSE(4, \"用户暂停\"),SYSTEM_SUSPENDED(5, \"系统暂停\")，TERMINATE(6, \"已终止\"),IN_QUEUE(7, \"排队中\"),SYSTEM_HANG_UP(10, \"系统挂起\"),WAITING_FOR_REDIAL(11, \"等待重呼\"),ACCOUNT_DISABLE(12, \"账户禁用\"),MAINTAIN(13, \"系统维护\");
	Status string `json:"status,omitempty"`
	// 系统挂起类型 ACCOUNT_DEBT(0, \"账户欠费\", \"使用的线路账户已欠费\"),NO_ROBOT_AVAILABLE(1, \"没有可用坐席\", \"当前没有可用坐席\"),PHONE_UNBIND(2, \"线路已解绑\", \"使用的线路已解绑\"),LINE_BREAKDOWN(3, \"线路故障\", \"使用的线路状态均为故障\"),AVAILABLE_ROBOT_NOT_ENOUGH(4,\"有效坐席数不足\",\"有效坐席数不足，请检查有效AI坐席数量！\");
	HangUpType string `json:"hangUpType,omitempty"`
	// 系统挂起原因
	HangUpReason string `json:"hangUpReason,omitempty"`
	// 已完成的任务总量
	CompletedTask int64 `json:"completedTask,omitempty"`
	// 任务总数
	TotalTask int64 `json:"totalTask,omitempty"`
	// 接通电话总量
	AnsweredTask int64 `json:"answeredTask,omitempty"`
	// AB类客户数量
	AblevelCustomer int64 `json:"ablevelCustomer,omitempty"`
	// 信息
	Msg string `json:"msg,omitempty"`
}