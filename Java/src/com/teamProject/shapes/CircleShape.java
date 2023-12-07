package com.teamProject.shapes;

import com.teamProject.Shape;

public class CircleShape implements Shape {
    private final int x;
    private final int y;
    private final int radius;
    private int currentX;
    private int currentY;

    public CircleShape(int x, int y, int radius) {
        this.radius = radius;
        this.x = x;
        this.y = y;
    }

    private double distance(int x1, int y1, int x2, int y2) {
        return Math.sqrt(Math.pow(x1 - x2, 2) + Math.pow(y1 - y2, 2));
    }

    @Override
    public int[] reset() {
        currentX = x - radius;
        currentY = y - radius;
        while (distance(currentX, currentY, x, y) > radius) {
            ++currentX;
        }
        return new int[]{currentX, currentY};
    }

    @Override
    public int[] next() {
        ++currentX;
        if (distance(currentX, currentY, x, y) > radius) {
            ++currentY;
            if (currentY >= y + radius) {
                return null;
            }
            currentX = x - radius;
            while (distance(currentX, currentY, x, y) > radius) {
                ++currentX;
            }
        }
        return new int[]{currentX, currentY};
    }
}
