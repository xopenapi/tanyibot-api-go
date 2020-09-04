# Go API client for tanyibot-api-go

tanyibot api

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./tanyibot-api-go"
```

## Documentation for API Endpoints

All URIs are relative to *https://openapi.tanyibot.com/apiOpen/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CallInterceptApi* | [**CallInterceptListGet**](docs/CallInterceptApi.md#callinterceptlistget) | **Get** /callIntercept/list | 获取拦截组列表接口
*CsRobotApi* | [**StaffGroupList**](docs/CsRobotApi.md#staffgrouplist) | **Get** /csRobot/staffGroupList | 获取人工坐席列表接口
*DialogFlowApi* | [**DialogFlowBatchUploadPost**](docs/DialogFlowApi.md#dialogflowbatchuploadpost) | **Post** /dialogFlow/batchUpload | 上传话术录音
*DialogFlowApi* | [**DialogFlowCopyDialogFlowPost**](docs/DialogFlowApi.md#dialogflowcopydialogflowpost) | **Post** /dialogFlow/copyDialogFlow | 话术复制
*DialogFlowApi* | [**DialogFlowGetDialogFlowCallJobRelatedInfoGet**](docs/DialogFlowApi.md#dialogflowgetdialogflowcalljobrelatedinfoget) | **Get** /dialogFlow/getDialogFlowCallJobRelatedInfo | 获取话术中存在人工介入和转人工等标识
*DialogFlowApi* | [**DialogFlowGetTotalDialogFlowListPost**](docs/DialogFlowApi.md#dialogflowgettotaldialogflowlistpost) | **Post** /dialogFlow/getTotalDialogFlowList | 话术列表
*DialogFlowApi* | [**ExportDialogFlowWordDoc**](docs/DialogFlowApi.md#exportdialogflowworddoc) | **Get** /dialogFlow/exportDialogFlowWordDoc | 获取话术主流程word文档
*DialogFlowApi* | [**GetDialogContentInfo**](docs/DialogFlowApi.md#getdialogcontentinfo) | **Get** /dialogFlow/getDialogContentInfo | 获取话术中所有需要录音的文本和录音文件
*DialogFlowApi* | [**GetDialogFlowList**](docs/DialogFlowApi.md#getdialogflowlist) | **Get** /dialogFlow/getDialogFlowList | 获取公司的机器人话术列表接口
*IntentLevelTagApi* | [**GetUsedIntentLevelTagList**](docs/IntentLevelTagApi.md#getusedintentleveltaglist) | **Get** /intentLevelTag/getUsedIntentLevelTagList | 获取已使用的意向标签组列表接口
*IntentLevelTagApi* | [**IntentLevelTagGetIntentLevelTagGet**](docs/IntentLevelTagApi.md#intentleveltaggetintentleveltagget) | **Get** /intentLevelTag/getIntentLevelTag | 获取意向标签内容
*IsvApi* | [**GenSubAccountIsv**](docs/IsvApi.md#gensubaccountisv) | **Post** /isv/genSubAccountIsv | 通过此接口可生成子账号的ISV账号
*IsvApi* | [**GetIsvList**](docs/IsvApi.md#getisvlist) | **Post** /isv/getIsvList | 获取isv列表
*IsvApi* | [**UpdateIsvInfo**](docs/IsvApi.md#updateisvinfo) | **Post** /isv/updateIsvInfo | 修改ISV对象的公司签名和回调地址
*JobApi* | [**Create**](docs/JobApi.md#create) | **Post** /job/create | 通过此接口可以创建新的任务
*JobApi* | [**ImportCustomer**](docs/JobApi.md#importcustomer) | **Post** /job/importCustomer | 
*JobApi* | [**JobCheckAllPost**](docs/JobApi.md#jobcheckallpost) | **Post** /job/checkAll | 
*JobApi* | [**JobDeletePost**](docs/JobApi.md#jobdeletepost) | **Post** /job/delete | 
*JobApi* | [**JobExecuteJobsPost**](docs/JobApi.md#jobexecutejobspost) | **Post** /job/executeJobs | 
*JobApi* | [**JobGetJobDetailGet**](docs/JobApi.md#jobgetjobdetailget) | **Get** /job/getJobDetail | 
*JobApi* | [**JobGetJobStatsInfoListPost**](docs/JobApi.md#jobgetjobstatsinfolistpost) | **Post** /job/getJobStatsInfoList | 
*JobApi* | [**JobGetJobsGet**](docs/JobApi.md#jobgetjobsget) | **Get** /job/getJobs | 获取任务列表接口
*JobApi* | [**JobModifyPatch**](docs/JobApi.md#jobmodifypatch) | **Patch** /job/modify | 修改任务接口
*JobApi* | [**JobPausePost**](docs/JobApi.md#jobpausepost) | **Post** /job/pause | 
*JobApi* | [**JobReAddCustomerToJobPost**](docs/JobApi.md#jobreaddcustomertojobpost) | **Post** /job/reAddCustomerToJob | 
*JobApi* | [**JobUpdateAiCountPost**](docs/JobApi.md#jobupdateaicountpost) | **Post** /job/updateAiCount | 
*JobApi* | [**Start**](docs/JobApi.md#start) | **Post** /job/start | 
*PhoneApi* | [**GetPhoneList**](docs/PhoneApi.md#getphonelist) | **Get** /phone/getPhoneList | 通过接口可以获取所有可用的外呼线路的列表
*PhoneApi* | [**UpdatePhoneInfoByTenantPhoneNumberId**](docs/PhoneApi.md#updatephoneinfobytenantphonenumberid) | **Post** /phone/updatePhoneInfoByTenantPhoneNumberId | 通过此接口修改线路的归属地、行业、黑名单，只能修改归属客户自己的线路
*PhoneApi* | [**UpdatePhonePriceByTenantPhoneNumberId**](docs/PhoneApi.md#updatephonepricebytenantphonenumberid) | **Post** /phone/updatePhonePriceByTenantPhoneNumberId | 修改绑定客户线路的价格
*PolicyGroupApi* | [**PolicyGroupCreatePost**](docs/PolicyGroupApi.md#policygroupcreatepost) | **Post** /policyGroup/create | 创建外呼策略组接口
*PolicyGroupApi* | [**PolicyGroupGetIdAndNamePairListGet**](docs/PolicyGroupApi.md#policygroupgetidandnamepairlistget) | **Get** /policyGroup/getIdAndNamePairList | 获取外呼策略组选择接口
*PolicyGroupApi* | [**PolicyGroupListGet**](docs/PolicyGroupApi.md#policygrouplistget) | **Get** /policyGroup/list | 获取外呼策略组列表接口
*PolicyGroupApi* | [**PolicyGroupUpdatePost**](docs/PolicyGroupApi.md#policygroupupdatepost) | **Post** /policyGroup/update | 更新外呼策略组接口
*TenantApi* | [**GetTenant**](docs/TenantApi.md#gettenant) | **Get** /tenant/getTenant | 获取公司列表接口


## Documentation For Models

 - [ApiResponse](docs/ApiResponse.md)
 - [BaseReq](docs/BaseReq.md)
 - [CallIntercept](docs/CallIntercept.md)
 - [CallPolicyGroupDetail](docs/CallPolicyGroupDetail.md)
 - [CheckAllRsp](docs/CheckAllRsp.md)
 - [CheckAllRspAllOf](docs/CheckAllRspAllOf.md)
 - [CopyDialogFlowReq](docs/CopyDialogFlowReq.md)
 - [CopyDialogFlowRsp](docs/CopyDialogFlowRsp.md)
 - [CopyDialogFlowRspAllOf](docs/CopyDialogFlowRspAllOf.md)
 - [CreatePolicyGroupReq](docs/CreatePolicyGroupReq.md)
 - [CustomerPerson](docs/CustomerPerson.md)
 - [DialogContentInfo](docs/DialogContentInfo.md)
 - [DialogContentInfoAllOf](docs/DialogContentInfoAllOf.md)
 - [DialogFlow](docs/DialogFlow.md)
 - [DialogFlowCallJobRelatedInfo](docs/DialogFlowCallJobRelatedInfo.md)
 - [DialogFlowCallJobRelatedInfoAllOf](docs/DialogFlowCallJobRelatedInfoAllOf.md)
 - [DialogFlowConfig](docs/DialogFlowConfig.md)
 - [DialogFlowList](docs/DialogFlowList.md)
 - [DialogFlowListAllOf](docs/DialogFlowListAllOf.md)
 - [DialogFlowNode](docs/DialogFlowNode.md)
 - [DialogFlowPage](docs/DialogFlowPage.md)
 - [DialogFlowStep](docs/DialogFlowStep.md)
 - [ExecuteJobInfo](docs/ExecuteJobInfo.md)
 - [ExecuteJobsReq](docs/ExecuteJobsReq.md)
 - [ExecuteJobsRsp](docs/ExecuteJobsRsp.md)
 - [ExecuteJobsRspAllOf](docs/ExecuteJobsRspAllOf.md)
 - [GetIdAndNamePairListRsp](docs/GetIdAndNamePairListRsp.md)
 - [GetIdAndNamePairListRspAllOf](docs/GetIdAndNamePairListRspAllOf.md)
 - [GetIntentLevelTagRsp](docs/GetIntentLevelTagRsp.md)
 - [GetIntentLevelTagRspAllOf](docs/GetIntentLevelTagRspAllOf.md)
 - [GetIsvListRsp](docs/GetIsvListRsp.md)
 - [GetIsvListRspAllOf](docs/GetIsvListRspAllOf.md)
 - [GetIsvListRspAllOfData](docs/GetIsvListRspAllOfData.md)
 - [GetJobDetailRsp](docs/GetJobDetailRsp.md)
 - [GetJobDetailRspAllOf](docs/GetJobDetailRspAllOf.md)
 - [GetJobStatsInfoListRsp](docs/GetJobStatsInfoListRsp.md)
 - [GetJobStatsInfoListRspAllOf](docs/GetJobStatsInfoListRspAllOf.md)
 - [GetJobStatsInfoListRspAllOfData](docs/GetJobStatsInfoListRspAllOfData.md)
 - [GetJobsRsp](docs/GetJobsRsp.md)
 - [GetJobsRspAllOf](docs/GetJobsRspAllOf.md)
 - [GetPhoneListRsp](docs/GetPhoneListRsp.md)
 - [GetPhoneListRspAllOf](docs/GetPhoneListRspAllOf.md)
 - [GetPolicyGroupListRsp](docs/GetPolicyGroupListRsp.md)
 - [GetPolicyGroupListRspAllOf](docs/GetPolicyGroupListRspAllOf.md)
 - [GetSubAccountIsvReq](docs/GetSubAccountIsvReq.md)
 - [GetSubAccountIsvResp](docs/GetSubAccountIsvResp.md)
 - [GetTenantRsp](docs/GetTenantRsp.md)
 - [GetTenantRspAllOf](docs/GetTenantRspAllOf.md)
 - [GetTotalDialogFlowListReq](docs/GetTotalDialogFlowListReq.md)
 - [GetUsedIntentLevelTagListRsp](docs/GetUsedIntentLevelTagListRsp.md)
 - [GetUsedIntentLevelTagListRspAllOf](docs/GetUsedIntentLevelTagListRspAllOf.md)
 - [ImportCustomerReq](docs/ImportCustomerReq.md)
 - [Industry](docs/Industry.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [InlineObject2](docs/InlineObject2.md)
 - [InlineObject3](docs/InlineObject3.md)
 - [InlineObject4](docs/InlineObject4.md)
 - [InlineObject5](docs/InlineObject5.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [IntentLevelTag](docs/IntentLevelTag.md)
 - [IntentLevelTagDetails](docs/IntentLevelTagDetails.md)
 - [InterceptListRsp](docs/InterceptListRsp.md)
 - [InterceptListRspAllOf](docs/InterceptListRspAllOf.md)
 - [Job](docs/Job.md)
 - [JobCreateRsp](docs/JobCreateRsp.md)
 - [JobUpdateRsp](docs/JobUpdateRsp.md)
 - [JobUpdateRspAllOf](docs/JobUpdateRspAllOf.md)
 - [JobUpdateRspAllOfData](docs/JobUpdateRspAllOfData.md)
 - [JobsPage](docs/JobsPage.md)
 - [JobsPageContent](docs/JobsPageContent.md)
 - [KnowledgeStep](docs/KnowledgeStep.md)
 - [LineStatus](docs/LineStatus.md)
 - [Phone](docs/Phone.md)
 - [PhoneDeadArea](docs/PhoneDeadArea.md)
 - [PolicyGroup](docs/PolicyGroup.md)
 - [PolicyGroupIdAndNamePair](docs/PolicyGroupIdAndNamePair.md)
 - [ReAddCustomerToJobReq](docs/ReAddCustomerToJobReq.md)
 - [RobotCallJob](docs/RobotCallJob.md)
 - [RobotCallJobInactiveTimeList](docs/RobotCallJobInactiveTimeList.md)
 - [RobotKnowledge](docs/RobotKnowledge.md)
 - [StaffGroupListRsp](docs/StaffGroupListRsp.md)
 - [StaffGroupListRspAllOf](docs/StaffGroupListRspAllOf.md)
 - [StatsInfo](docs/StatsInfo.md)
 - [TotalDialogFlowListRsp](docs/TotalDialogFlowListRsp.md)
 - [TotalDialogFlowListRspAllOf](docs/TotalDialogFlowListRspAllOf.md)
 - [UpdateIsvInfoReq](docs/UpdateIsvInfoReq.md)
 - [UpdatePhoneInfoByTenantPhoneNumberIdReq](docs/UpdatePhoneInfoByTenantPhoneNumberIdReq.md)
 - [UpdatePhoneInfoByTenantPhoneNumberIdReqDeadArea](docs/UpdatePhoneInfoByTenantPhoneNumberIdReqDeadArea.md)
 - [UpdatePolicyGroupReq](docs/UpdatePolicyGroupReq.md)
 - [WechatPushCondition](docs/WechatPushCondition.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



