package com.teamProject.processors;

import com.teamProject.ImageProcessor;

public class BlurProcessor extends ImageProcessor {
    static final int R = 0;
    static final int G = 1;
    static final int B = 2;

    int boxSize;

    public BlurProcessor(int boxSize) {
        this.boxSize = boxSize;
        if (boxSize % 2 == 0) {
            ++this.boxSize;
        }
        this.boxSize = Math.min(this.boxSize, 4095);
    }

    private int toIndex(int x, int y) {
        return x + width * y;
    }

    private int[] fromPixelValue(int pixelValue) {
        return new int[]{(pixelValue >> 16) & 255, (pixelValue >> 8) & 255, pixelValue & 255};
    }

    private int toPixelValue(int[] channels) {
        return 0xFF000000 + channels[R] * 256 * 256 + channels[G] * 256 + channels[B];
    }

    @Override
    public int[] process(int[] pixels, int width, int height) {
        super.process(pixels, width, height);

        int boxArea = boxSize * boxSize;

        int[] pxBuf;
        int[] newPixels = new int[pixels.length * 3];

        int[] avgRow = new int[width * 3];
        int halfBoxSize = boxSize / 2;
        int halfBoxSize3 = halfBoxSize * 3;

        for (int i = 0; i < width; ++i) {
            for (int j = 0; j <= halfBoxSize; ++j) {
                pxBuf = fromPixelValue(pixels[toIndex(i, j)]);
                int i3 = i * 3;
                for (int k = 0; k < 3; ++k) {
                    avgRow[i3 + k] += pxBuf[k];
                }
            }
        }

        //first pixel
        for (int i = 0; i <= halfBoxSize; ++i) {
            int i3 = i * 3;
            for (int j = 0; j < 3; ++j) {
                newPixels[j] += avgRow[i3 + j];
            }
        }

        //other pixels in first row
        for (int i = 1; i < width; ++i) {
            int i3 = i * 3;
            for (int j = 0; j < 3; ++j) {
                int index = i3 + j;
                newPixels[index] = newPixels[index - 3];
                if (i < width - halfBoxSize) {
                    newPixels[index] += avgRow[index + halfBoxSize3];
                }
                if (i > halfBoxSize) {
                    newPixels[index] -= avgRow[index - halfBoxSize3 - 3];
                }
            }
        }

        //other rows
        int widthTimesHalfBoxSize = width * halfBoxSize;
        int temp1 = widthTimesHalfBoxSize;
        int temp2 = -width * (halfBoxSize + 1);
        for (int row = 1; row < height; ++row) {
            int widthTimesRow = width * row;
            temp1 += width;
            temp2 += width;
            int index = toIndex(0, row);
            int previousIndex = toIndex(0, row - 1);
            int[] tempAdd = new int[3];
            int[] tempSubtract = new int[3];
            int channelIndex = index * 3;
            int previousChannelIndex = previousIndex * 3;
            for (int i = 0; i < 3; ++i) {
                newPixels[channelIndex] = newPixels[previousChannelIndex];

                if (row < height - halfBoxSize) {
                    for (int j = 0; j <= halfBoxSize; ++j) {
                        tempAdd[i] += fromPixelValue(pixels[j + temp1])[i];
                    }

                    newPixels[channelIndex] += tempAdd[i];
                }

                if (row > halfBoxSize) {
                    for (int j = 0; j <= halfBoxSize; ++j) {
                        tempSubtract[i] += fromPixelValue(pixels[j + temp2])[i];
                    }

                    newPixels[channelIndex] -= tempSubtract[i];
                }
                ++channelIndex;
                ++previousChannelIndex;
            }

            int width3 = width * 3;
            int tempAddAddIndex = halfBoxSize + widthTimesRow + widthTimesHalfBoxSize;
            int tempAddSubtractIndex = -halfBoxSize - 1 + widthTimesRow + widthTimesHalfBoxSize;
            int tempSubtractSubtractIndex = -halfBoxSize - 1 + widthTimesRow - widthTimesHalfBoxSize - width;
            int tempSubtractAddIndex = halfBoxSize + widthTimesRow - widthTimesHalfBoxSize - width;
            for (int col = 1; col < width; ++col) {
                ++tempAddAddIndex;
                ++tempAddSubtractIndex;
                ++tempSubtractSubtractIndex;
                ++tempSubtractAddIndex;
                int[] tempAddAddPixel = new int[0];
                int[] tempAddSubtractPixel = new int[0];
                int[] tempSubtractSubtractPixel = new int[0];
                int[] tempSubtractAddPixel = new int[0];
                if (row < height - halfBoxSize) {
                    if (col < width - halfBoxSize) {
                        tempAddAddPixel = fromPixelValue(pixels[tempAddAddIndex]);
                    }
                    if (col > halfBoxSize) {
                        tempAddSubtractPixel = fromPixelValue(pixels[tempAddSubtractIndex]);
                    }
                }
                if (row > halfBoxSize) {
                    if (col > halfBoxSize) {
                        tempSubtractSubtractPixel = fromPixelValue(pixels[tempSubtractSubtractIndex]);
                    }
                    if (col < width - halfBoxSize) {
                        tempSubtractAddPixel = fromPixelValue(pixels[tempSubtractAddIndex]);
                    }
                }

                int columnIndex = (col + width * row) * 3;
                for (int i = 0; i < 3; ++i) {
                    newPixels[columnIndex] = newPixels[columnIndex - width3];
                    if (row < height - halfBoxSize) {
                        if (col < width - halfBoxSize) {
                            tempAdd[i] += tempAddAddPixel[i];
                        }
                        if (col > halfBoxSize) {
                            tempAdd[i] -= tempAddSubtractPixel[i];
                        }
                        newPixels[columnIndex] += tempAdd[i];
                    }
                    if (row > halfBoxSize) {
                        if (col > halfBoxSize) {
                            tempSubtract[i] -= tempSubtractSubtractPixel[i];
                        }
                        if (col < width - halfBoxSize) {
                            tempSubtract[i] += tempSubtractAddPixel[i];
                        }
                        newPixels[columnIndex] -= tempSubtract[i];
                    }
                    ++columnIndex;
                }
            }
        }

        int[] outPixels = new int[pixels.length];

        for (int i = 0; i < pixels.length; ++i) {
            outPixels[i] = toPixelValue(new int[]{newPixels[i * 3 + R] / boxArea, newPixels[i * 3 + G] / boxArea, newPixels[i * 3 + B] / boxArea});
        }

        return outPixels;
    }

    @Override
    public String getName() {
        return "Box blur (" + boxSize + "x" + boxSize + ")";
    }
}
