# -*- coding: utf-8 -*-

from PIL import Image
import cv2
import logging
import numpy as np
import os


class ImageStyleCover():

    def __init__(self, path=None):
        if path is not None:
            self.path = path
            self.img = Image.open(path)

    def setImageFromArray(self, array):
        self.img = Image.fromarray(array)
        return self

    def setImage(self, image):
        self.img = image
        return self

    def toCV2(self, image):
        return cv2.cvtColor(np.asarray(image), cv2.COLOR_RGB2BGR)

    def toImage(self, cv):
        return Image.fromarray(cv2.cvtColor(img, cv2.COLOR_BGR2RGB))
    
    def toSketch(self, type='image'):
        L = np.asarray(self.img.convert('L')).astype('float')

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

        gd = gd.astype('uint8')

        if type == 'image':
            return Image.fromarray(gd)  # 重构图像
        else:
            return gd

    def toSumiao(self):
        img_rgb = self.toCV2(self.img)
        img_gray = cv2.cvtColor(img_rgb, cv2.COLOR_RGB2GRAY)
        img_blur = cv2.GaussianBlur(img_gray, ksize=(21, 21),
                            sigmaX=0, sigmaY=0)
        return cv2.divide(img_gray, img_blur, scale=255)

    def toCartoon(self):
        num_down = 2  # 缩减像素采样的数目
        num_bilateral = 7  # 定义双边滤波的数目
        img_rgb = self.toCV2(self.img)
        # img_rgb = cv2.imread(imgInput_FileName)     #读取图片
        # 用高斯金字塔降低取样
        img_color = img_rgb

        # 重复使用小的双边滤波代替一个大的滤波
        for _ in range(num_bilateral):
            img_color = cv2.bilateralFilter(
                img_color, d=9, sigmaColor=9, sigmaSpace=7)

        # 转换为灰度并且使其产生中等的模糊
        img_gray = cv2.cvtColor(img_rgb, cv2.COLOR_RGB2GRAY)
        img_blur = cv2.medianBlur(img_gray, 5)
        # 检测到边缘并且增强其效果
        img_edge = cv2.adaptiveThreshold(img_blur, 255,
                                         cv2.ADAPTIVE_THRESH_MEAN_C,
                                         cv2.THRESH_BINARY,
                                         blockSize=9,
                                         C=2)
        # 转换回彩色图像
        img_edge = cv2.cvtColor(img_edge, cv2.COLOR_GRAY2RGB)
        return cv2.bitwise_and(img_color, img_edge)


if __name__ == '__main__':
    cover = ImageStyleCover("../images/tupian.jpg")
    img = cover.toSketch()
    img.save("../images/tupian_cover1.jpg")
    img = cover.toSumiao()
    cv2.imwrite('../images/tupian_sumiao.jpeg', img)
    img = cover.toCartoon()
    cv2.imwrite("../images/tupian_cartoon.jpg", img)
    cover = ImageStyleCover("../images/p.jpeg")
    img = cover.toSumiao()
    cv2.imwrite('../images/p_sumiao.jpeg', img)
    img = cover.toCartoon()
    cv2.imwrite("../images/p_cartoon.jpg", img)

