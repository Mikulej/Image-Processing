package com.teamProject;

public abstract class ImageProcessor {
    protected int width;
    protected int height;

    public int[] process(int[] pixels, int width, int height) {
        this.width = width;
        this.height = height;
        return pixels;
    }

    public ProcessResult test(int[] pixels, int width, int height) {
        long startTime = System.nanoTime();
        int[] newPixels = process(pixels, width, height);
        double elapsedTime = (System.nanoTime() - startTime) / 1000000000.0;
        return new ProcessResult(newPixels, elapsedTime);
    }

    public abstract String getName();
}
