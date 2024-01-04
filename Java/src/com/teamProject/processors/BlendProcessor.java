package com.teamProject.processors;

import com.teamProject.ImageProcessor;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

public class BlendProcessor extends ImageProcessor {
    static final int R = 0;
    static final int G = 1;
    static final int B = 2;

    String filename;
    BufferedImage blendImg = null;
    int blendWidth;
    int blendHeight;
    int[] blendPixels;
    double amount;

    public BlendProcessor(String filename, double amount) {
        this.filename = filename + ".jpg";
        try {
            blendImg = ImageIO.read(new File("images/" + this.filename));
        } catch (IOException e) {
            e.printStackTrace();
        }
        blendPixels = blendImg.getRGB(0, 0, blendImg.getWidth(), blendImg.getHeight(), null, 0, blendImg.getWidth());
        blendWidth = blendImg.getWidth();
        blendHeight = blendImg.getHeight();

        this.amount = amount;
    }

    private int toIndex(int x, int y) {
        return x + width * y;
    }

    private int toBlendIndex(int x, int y) {
        return x + blendWidth * y;
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

        int[] newPixels = new int[pixels.length];

        for (int y = 0; y < height; ++y) {
            for (int x = 0; x < width; ++x) {
                int[] channels = fromPixelValue(pixels[toIndex(x, y)]);
                int[] blendChannels = {0, 0, 0};
                if (x < blendImg.getWidth() && y < blendImg.getHeight()) {
                    blendChannels = fromPixelValue(blendPixels[toBlendIndex(x, y)]);
                }
                int[] newChannels = new int[3];

                for (int channel = 0; channel < 3; ++channel) {
                    newChannels[channel] = (int) (amount * blendChannels[channel] + (1 - amount) * channels[channel]);
                    if (newChannels[channel] > 255) newChannels[channel] = 255;
                }

                newPixels[toIndex(x, y)] = toPixelValue(newChannels);
            }
        }

        return newPixels;
    }

    @Override
    public String getName() {
        return "Blend (" + filename.replace("/", "(slash)") + ")";
    }
}
