package v1

import (
	"github.com/1Panel-dev/1Panel/backend/app/api/v1/helper"
	"github.com/1Panel-dev/1Panel/backend/app/dto"
	"github.com/1Panel-dev/1Panel/backend/constant"
	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// @Tags Container
// @Summary Page containers
// @Description 获取容器列表分页
// @Accept json
// @Param request body dto.PageContainer true "request"
// @Produce json
// @Success 200 {object} dto.PageResult
// @Security ApiKeyAuth
// @Router /containers/search [post]
func (b *BaseApi) SearchContainer(c *gin.Context) {
	var req dto.PageContainer
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	total, list, err := containerService.Page(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, dto.PageResult{
		Items: list,
		Total: total,
	})
}

// @Tags Container Compose
// @Summary Page composes
// @Description 获取编排列表分页
// @Accept json
// @Param request body dto.SearchWithPage true "request"
// @Success 200 {object} dto.PageResult
// @Security ApiKeyAuth
// @Router /containers/compose/search [post]
func (b *BaseApi) SearchCompose(c *gin.Context) {
	var req dto.SearchWithPage
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	total, list, err := containerService.PageCompose(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, dto.PageResult{
		Items: list,
		Total: total,
	})
}

// @Tags Container Compose
// @Summary Test compose
// @Description 测试 compose 是否可用
// @Accept json
// @Param request body dto.ComposeCreate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/compose/test [post]
// @x-panel-log {"bodyKeys":["name"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"检测 compose [name] 格式","formatEN":"check compose [name]"}
func (b *BaseApi) TestCompose(c *gin.Context) {
	var req dto.ComposeCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	isOK, err := containerService.TestCompose(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, isOK)
}

// @Tags Container Compose
// @Summary Create compose
// @Description 创建容器编排
// @Accept json
// @Param request body dto.ComposeCreate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/compose [post]
// @x-panel-log {"bodyKeys":["name"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"创建 compose [name]","formatEN":"create compose [name]"}
func (b *BaseApi) CreateCompose(c *gin.Context) {
	var req dto.ComposeCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	log, err := containerService.CreateCompose(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, log)
}

// @Tags Container Compose
// @Summary Operate compose
// @Description 容器编排操作
// @Accept json
// @Param request body dto.ComposeOperation true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/compose/operate [post]
// @x-panel-log {"bodyKeys":["name","operation"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"compose [operation] [name]","formatEN":"compose [operation] [name]"}
func (b *BaseApi) OperatorCompose(c *gin.Context) {
	var req dto.ComposeOperation
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	if err := containerService.ComposeOperation(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Container
// @Summary Update container
// @Description 更新容器
// @Accept json
// @Param request body dto.ContainerOperate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/update [post]
// @x-panel-log {"bodyKeys":["name","image"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"更新容器 [name][image]","formatEN":"update container [name][image]"}
func (b *BaseApi) ContainerUpdate(c *gin.Context) {
	var req dto.ContainerOperate
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := containerService.ContainerUpdate(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Container
// @Summary Load container info
// @Description 获取容器表单信息
// @Accept json
// @Param request body dto.OperationWithName true "request"
// @Success 200 {object} dto.ContainerOperate
// @Security ApiKeyAuth
// @Router /containers/info [post]
func (b *BaseApi) ContainerInfo(c *gin.Context) {
	var req dto.OperationWithName
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	data, err := containerService.ContainerInfo(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, data)
}

// @Summary Load container limis
// @Description 获取容器限制
// @Success 200 {object} dto.ResourceLimit
// @Security ApiKeyAuth
// @Router /containers/limit [get]
func (b *BaseApi) LoadResouceLimit(c *gin.Context) {
	data, err := containerService.LoadResouceLimit()
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, data)
}

// @Tags Container
// @Summary Create container
// @Description 创建容器
// @Accept json
// @Param request body dto.ContainerOperate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers [post]
// @x-panel-log {"bodyKeys":["name","image"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"创建容器 [name][image]","formatEN":"create container [name][image]"}
func (b *BaseApi) ContainerCreate(c *gin.Context) {
	var req dto.ContainerOperate
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := containerService.ContainerCreate(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Container
// @Summary Clean container
// @Description 容器清理
// @Accept json
// @Param request body dto.ContainerPrune true "request"
// @Success 200 {object} dto.ContainerPruneReport
// @Security ApiKeyAuth
// @Router /containers/prune [post]
// @x-panel-log {"bodyKeys":["pruneType"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"清理容器 [pruneType]","formatEN":"clean container [pruneType]"}
func (b *BaseApi) ContainerPrune(c *gin.Context) {
	var req dto.ContainerPrune
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	report, err := containerService.Prune(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, report)
}

// @Tags Container
// @Summary Clean container log
// @Description 清理容器日志
// @Accept json
// @Param request body dto.OperationWithName true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/clean/log [post]
// @x-panel-log {"bodyKeys":["name"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"清理容器 [name] 日志","formatEN":"clean container [name] logs"}
func (b *BaseApi) CleanContainerLog(c *gin.Context) {
	var req dto.OperationWithName
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := containerService.ContainerLogClean(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Container
// @Summary Operate Container
// @Description 容器操作
// @Accept json
// @Param request body dto.ContainerOperation true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/operate [post]
// @x-panel-log {"bodyKeys":["name","operation","newName"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"容器 [name] 执行 [operation] [newName]","formatEN":"container [operation] [name] [newName]"}
func (b *BaseApi) ContainerOperation(c *gin.Context) {
	var req dto.ContainerOperation
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := containerService.ContainerOperation(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Container
// @Summary Container stats
// @Description 容器监控信息
// @Param id path integer true "容器id"
// @Success 200 {object} dto.ContainterStats
// @Security ApiKeyAuth
// @Router /containers/stats/:id [get]
func (b *BaseApi) ContainerStats(c *gin.Context) {
	containerID, ok := c.Params.Get("id")
	if !ok {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, errors.New("error container id in path"))
		return
	}

	result, err := containerService.ContainerStats(containerID)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, result)
}

// @Tags Container
// @Summary Container inspect
// @Description 容器详情
// @Accept json
// @Param request body dto.InspectReq true "request"
// @Success 200 {string} result
// @Security ApiKeyAuth
// @Router /containers/inspect [post]
func (b *BaseApi) Inspect(c *gin.Context) {
	var req dto.InspectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	result, err := containerService.Inspect(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, result)
}

// @Tags Container
// @Summary Container logs
// @Description 容器日志
// @Param container query string false "容器名称"
// @Param since query string false "时间筛选"
// @Param follow query string false "是否追踪"
// @Param tail query string false "显示行号"
// @Security ApiKeyAuth
// @Router /containers/search/log [post]
func (b *BaseApi) ContainerLogs(c *gin.Context) {
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.LOG.Errorf("gin context http handler failed, err: %v", err)
		return
	}
	defer wsConn.Close()

	container := c.Query("container")
	since := c.Query("since")
	follow := c.Query("follow") == "true"
	tail := c.Query("tail")

	if err := containerService.ContainerLogs(wsConn, container, since, tail, follow); err != nil {
		_ = wsConn.WriteMessage(1, []byte(err.Error()))
		return
	}
}

// @Tags Container Network
// @Summary Page networks
// @Description 获取容器网络列表分页
// @Accept json
// @Param request body dto.SearchWithPage true "request"
// @Produce json
// @Success 200 {object} dto.PageResult
// @Security ApiKeyAuth
// @Router /containers/network/search [post]
func (b *BaseApi) SearchNetwork(c *gin.Context) {
	var req dto.SearchWithPage
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	total, list, err := containerService.PageNetwork(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, dto.PageResult{
		Items: list,
		Total: total,
	})
}

// @Tags Container Network
// @Summary Delete network
// @Description 删除容器网络
// @Accept json
// @Param request body dto.BatchDelete true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/network/del [post]
// @x-panel-log {"bodyKeys":["names"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"删除容器网络 [names]","formatEN":"delete container network [names]"}
func (b *BaseApi) DeleteNetwork(c *gin.Context) {
	var req dto.BatchDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	if err := containerService.DeleteNetwork(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Container Network
// @Summary Create network
// @Description 创建容器网络
// @Accept json
// @Param request body dto.NetworkCreate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/network [post]
// @x-panel-log {"bodyKeys":["name"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"创建容器网络 name","formatEN":"create container network [name]"}
func (b *BaseApi) CreateNetwork(c *gin.Context) {
	var req dto.NetworkCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	if err := containerService.CreateNetwork(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Container Volume
// @Summary Page volumes
// @Description 获取容器存储卷分页
// @Accept json
// @Param request body dto.SearchWithPage true "request"
// @Produce json
// @Success 200 {object} dto.PageResult
// @Security ApiKeyAuth
// @Router /containers/volume/search [post]
func (b *BaseApi) SearchVolume(c *gin.Context) {
	var req dto.SearchWithPage
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	total, list, err := containerService.PageVolume(req)
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, dto.PageResult{
		Items: list,
		Total: total,
	})
}

// @Tags Container Volume
// @Summary List volumes
// @Description 获取容器存储卷列表
// @Accept json
// @Param request body dto.PageInfo true "request"
// @Produce json
// @Success 200 {object} dto.PageResult
// @Security ApiKeyAuth
// @Router /containers/volume/search [get]
func (b *BaseApi) ListVolume(c *gin.Context) {
	list, err := containerService.ListVolume()
	if err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, list)
}

// @Tags Container Volume
// @Summary Delete volume
// @Description 删除容器存储卷
// @Accept json
// @Param request body dto.BatchDelete true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/volume/del [post]
// @x-panel-log {"bodyKeys":["names"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"删除容器存储卷 [names]","formatEN":"delete container volume [names]"}
func (b *BaseApi) DeleteVolume(c *gin.Context) {
	var req dto.BatchDelete
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	if err := containerService.DeleteVolume(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Container Volume
// @Summary Create volume
// @Description 创建容器存储卷
// @Accept json
// @Param request body dto.VolumeCreate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/volume [post]
// @x-panel-log {"bodyKeys":["name"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"创建容器存储卷 [name]","formatEN":"create container volume [name]"}
func (b *BaseApi) CreateVolume(c *gin.Context) {
	var req dto.VolumeCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	if err := containerService.CreateVolume(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}

// @Tags Container Compose
// @Summary Update compose
// @Description 更新容器编排
// @Accept json
// @Param request body dto.ComposeUpdate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /containers/compose/update [post]
// @x-panel-log {"bodyKeys":["name"],"paramKeys":[],"BeforeFuntions":[],"formatZH":"更新 compose [name]","formatEN":"update compose information [name]"}
func (b *BaseApi) ComposeUpdate(c *gin.Context) {
	var req dto.ComposeUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}
	if err := global.VALID.Struct(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrBadRequest, constant.ErrTypeInvalidParams, err)
		return
	}

	if err := containerService.ComposeUpdate(req); err != nil {
		helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
		return
	}
	helper.SuccessWithData(c, nil)
}
