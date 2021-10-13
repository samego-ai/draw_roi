<h1 align="center">
    <a href="https://github.com/samego-ai/draw_roi">
        点线框绘制组件
    </a>
</h1>
<p align="center">
    基于 GRPC 通信方式绘制 ROI 的组件服务
     <br>
    更加优雅的服务姿势、更加灵活的独立服务，让应用层服务组件更加标准规范
</p>
<p align="center">
    <a href="https://goreportcard.com/report/github.com/samego-ai/draw_roi">
        <img src="https://goreportcard.com/badge/github.com/samego-ai/draw_roi" alt="Go Report">
    </a>
    <a href="https://packagist.org/packages/alicfeng/aliyun_rocket_mq">
        <img src="https://img.shields.io/docker/pulls/samego/draw_roi.svg" alt="Docker Pull">
    </a>
    <a href="https://packagist.org/packages/alicfeng/aliyun_rocket_mq">
        <img src="https://poser.pugx.org/alicfeng/aliyun_rocket_mq/license.svg" alt="License">
    </a>
</p>




## 应用场景

算法分析识别到目标检测物，需要在对应帧或图片绘制对应的框。



## 功能支持

- [x] 多框绘制 且 文本绘写



## 配置说明

|        字段         |           说明           |                  示例                  |
| :-----------------: | :----------------------: | :------------------------------------: |
|  `SG_SERVER_PROTOCOL` | 服务通讯协议，支持 `tcp` |                 `tcp`                  |
|  `SG_SERVER_ADDRESS` |       服务通讯地址       |             `0.0.0.0:1280`             |
| `SG_DRAW_FONT_PATH` |     绘制笔刷字体路径     | `/usr/local/draw_roi/fonts/default.ttf` |
| `SG_LOGGER_LEVEL` |     日志级别     | 基于 `kataras/golog` |
| `SG_LOGGER_PREFIX` |    日志前缀     | `samego.ai.draw_roi` |



## 服务部署

```shell
docker run -itd -p 1280:1280 --name draw_roi samego/draw_roi:latest
```



## 快速使用

详细请查看 `pb/*.proto` 文件

1. 绘制多框

   ```protobuf
   syntax = "proto3";
   package pb;
   
   service Draw {
   rpc MultiRectangle (MultiRectangleRequest) returns (MultiRectangleResponse) {}
   }
   
   message MultiRectangleRequest {
   message rectangle  {
      int32 X = 1;
      int32 Y = 2;
      int32 W = 3;
      int32 H = 4;
   }
      string Image = 1;
      string Title = 2;
      repeated rectangle Rectangle = 3;
   }
   message MultiRectangleResponse {
   message data {
    string Image = 1;
   }
      int32 Code = 1;
      string Message = 2;
      data Data = 3;
   }
   ```

   





