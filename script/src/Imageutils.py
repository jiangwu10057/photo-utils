# -*- coding: utf-8 -*-

import cv2
import logging
import numpy as np
import os
from PIL import Image
import argparse
from enum import Enum

class CoverStyle(Enum):
    CARTOON = 2
    GRAY = 1
    SKETCH = 3
    RECOVER = -1


class ImageCoverUtils():

    def __init__(self, path=None):
        self.setPath(path)

    def setPath(self, path):
        if path is not None:
            self.path = path
            self.image = cv2.imread(self.path)
        return self

    def setImageFromArray(self, image):
        self.image = cv2.cvtColor(image, cv2.COLOR_RGB2BGR)
        return self

    def setTarget(self, target):
        self.target = target
        return self

    def toGray(self):
        return cv2.cvtColor(self.image, cv2.COLOR_RGB2GRAY)

    def toSketch(self):
        img_gray = self.toGray()
        img_blur = cv2.GaussianBlur(img_gray, ksize=(21, 21),
                                    sigmaX=0, sigmaY=0)
        return cv2.divide(img_gray, img_blur, scale=255)

    def toCartoon(self):
        # 用高斯金字塔降低取样
        img_color = self.image

        num_bilateral = 7  # 定义双边滤波的数目

        # 重复使用小的双边滤波代替一个大的滤波
        for _ in range(num_bilateral):
            img_color = cv2.bilateralFilter(
                img_color, d=9, sigmaColor=9, sigmaSpace=7)

        # 转换为灰度并且使其产生中等的模糊
        img_gray = self.toGray()
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

    def write(self, image):
        cv2.imwrite(self.target, image)

    def recover(self):
        height, width, depth = self.image[0:3]

        thresh = cv2.inRange(self.image, np.array([240, 240, 240]), np.array([255, 255, 255]))

        kernel = np.ones((3, 3), np.uint8)

        hi_mash = cv2.dilate(thresh, kernel, iterations=1)
        specular = cv2.inpaint(self.image, hi_mash, 5, flags=cv2.INPAINT_NS)

        self.write(specular)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='manual to this script')
    parser.add_argument("--src", type=str, default="")
    parser.add_argument("--dist", type=str, default="")
    parser.add_argument("--type", type=int, default=CoverStyle.GRAY)
    args = parser.parse_args()

    if args.src == "":
        print(False)
    else:
        if args.dist == "":
            args.dist = args.src
        
        cover = ImageCoverUtils(args.src)

        if args.type == CoverStyle.SKETCH.value:
            img = cover.toSketch()
            cover.setTarget(args.dist).write(img)
        elif args.type == CoverStyle.GRAY.value:
            img = cover.toGray()
            cover.setTarget(args.dist).write(img)
        elif args.type == CoverStyle.CARTOON.value:
            img = cover.toCartoon()
            cover.setTarget(args.dist).write(img)
        elif args.type == CoverStyle.RECOVER.value:
            cover.setTarget(args.dist).recover()
        else:
            print(False)
            exit(0)
    
    print(True)
    exit(0)


    # cover = ImageCoverUtils("../images/tupian.jpg")
    # img = cover.toSketch()
    # cover.setTarget('../images/tupian_sketch.jpeg').write(img)
    # img = cover.toGray()
    # cover.setTarget('../images/tupian_gray.jpeg').write(img)
    # img = cover.toCartoon()
    # cover.setTarget('../images/tupian_cartoon.jpeg').write(img)