# -*- coding: utf-8 -*-

from Imageutils import ImageCoverUtils
from moviepy.editor import VideoFileClip
import logging
import os
import cv2
import numpy as np
from PIL import Image
from enum import Enum

class CoverStyle(Enum):
    CARTOON = 'cartoon'
    GRAY = 'gray'
    SKETCH = 'sketch'

class VedioStyleCover():

    def __init__(self, path=None):
        self.imageCoverUtils = ImageCoverUtils()
        self.setPath(path)

    def setPath(self, path):
        self.path = path
        self.vido = VideoFileClip(self.path)
        return self

    def setTarget(self, target):
        self.target = target
        return self

    def setStyle(self, style):
        self.style = style
        return self

    def getBgMusic(self):
        return self.video.audio

    def addBgMusic(self, audio):
        video = VideoFileClip(video_name)  # 设置视频的音频
        video = video.set_audio(audio)  # 保存新的视频文件
        video.write_videofile(output_video)

    def _toThreeChannel(self, array):
        image = np.expand_dims(array, axis = 2)
        return np.concatenate((image,image,image),axis = -1)

    def process_image(self, frame):
        cover = self.imageCoverUtils.setImageFromArray(frame)
        if self.style == CoverStyle.CARTOON:
            return cover.toCartoon()
        elif self.style == CoverStyle.GRAY:
            return self._toThreeChannel(cover.toGray())
        elif self.style == CoverStyle.SKETCH:
            return self._toThreeChannel(cover.toSketch())
        else:
            return frame

    def cover(self):
        myClip = VideoFileClip(self.path)
        result = myClip.fl_image(self.process_image)
        result = result.set_audio(myClip.audio)
        result.write_videofile(self.target)

    def toSketch(self, frame):
        image = Image.fromarray(frame)
        L = np.asarray(image.convert('L')).astype('float')
        depth = 10.  # (0-100)
        grad = np.gradient(L)  # 取图像灰度的梯度值
        grad_x, grad_y = grad  # 分别取横纵图像梯度值
        grad_x = grad_x * depth / 100.
        grad_y = grad_y * depth / 100.
        A = np.sqrt(grad_x ** 2 + grad_y ** 2 + 1.)
        uni_x = grad_x / A
        uni_y = grad_y / A
        uni_z = 1. / A

        el = np.pi / 2.2  # 光源的俯视角度，弧度值
        az = np.pi / 4  # 光源的方位角度，弧度值
        dx = np.cos(el) * np.cos(az)  # 光源对x轴的影响
        dy = np.cos(el) * np.sin(az)  # 光源对y轴的影响
        dz = np.sin(el)  # 光源对z轴的影响

        gd = 255 * (dx * uni_x + dy * uni_y + dz * uni_z)  # 光源归一化
        gd = gd.clip(0, 255)  # 避免数据越界，将生成的灰度值裁剪至0-255之间

        return Image.fromarray(gd.astype('uint8'))  # 重构图像

    def dealVedio(self):
        if self.target is None:
            return

        videoCapture = cv2.VideoCapture(self.path)

        fps = videoCapture.get(cv2.CAP_PROP_FPS)  # 获取帧率
        size = (int(videoCapture.get(cv2.CAP_PROP_FRAME_WIDTH)),  # 获取视频尺寸
                int(videoCapture.get(cv2.CAP_PROP_FRAME_HEIGHT)))

        if videoCapture.isOpened():
            rval,frame = videoCapture.read()
        else:
            rval = False

        video_writer = cv2.VideoWriter(self.target, cv2.VideoWriter_fourcc(*'mp4v'),fps, size)

        while rval:
            rval,frame = videoCapture.read()# 读取视频帧
            img = self.toSketch(frame)
            frame_converted = np.array(img)

            # 转化为三通道
            image = np.expand_dims(frame_converted,axis = 2)
            result_arr = np.concatenate((image,image,image),axis = -1)
            result_arr = np.hstack((frame,result_arr))

            video_writer.write(result_arr)

        video_writer.release()

if __name__ == '__main__':
    vedioStyleCover = VedioStyleCover("../images/20210210-084926-928.mp4")
    vedioStyleCover.setTarget('../images/gray.mp4').setStyle(CoverStyle.GRAY).cover()
    vedioStyleCover.setTarget('../images/cartoon.mp4').setStyle(CoverStyle.CARTOON).cover()
    vedioStyleCover.setTarget('../images/sketch.mp4').setStyle(CoverStyle.SKETCH).cover()