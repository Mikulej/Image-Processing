package com.teamProject;

import com.teamProject.processors.*;
import com.teamProject.shapes.*;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

public class Main {
    static final int TESTS_NUM = 12;
    static final String[] FILENAMES = {
            "1k/rock_wall_10_diff_1k",
            "1k/wood_planks_diff_1k",
            "2k/rock_wall_10_diff_2k",
            "2k/wood_planks_diff_2k",
            "4k/rock_wall_10_diff_4k",
            "4k/wood_planks_diff_4k"
    };
    static final ImageProcessor[] PROCESSORS = {
            new BlurProcessor(11),
            new BlendProcessor("blend", 0.1),
            new RemoveProcessor(new Shape[]{
                    new RectangleShape(128, 128, 256, 256),
                    new RectangleShape(576, 640, 384, 256),
                    new CircleShape(768, 256, 128),
                    new CircleShape(256, 768, 64)
            }, new int[]{0, 0, 0})
    };

    public static void main(String[] args) {
        for (ImageProcessor processor : PROCESSORS) {
            System.out.printf("Testing algorithm: %s%n", processor.getName());
            for (String filename : FILENAMES) {
                try {
                    BufferedImage img = ImageIO.read(new File("images/" + filename + ".jpg"));
                    System.out.printf("\tLoaded image \"images/%s.jpg\" (%dx%d) for %s%n", filename, img.getWidth(), img.getHeight(), processor.getName());

                    double averageTime = 0.0;
                    for (int i = 0; i < TESTS_NUM; ++i) {
                        System.out.printf("\t\tAttempt %d: ", i + 1);

                        int[] pixels = img.getRGB(0, 0, img.getWidth(), img.getHeight(), null, 0, img.getWidth());
                        ProcessResult result = processor.test(pixels, img.getWidth(), img.getHeight());
                        System.out.printf("Time: %g s%n", result.elapsedTime);
                        averageTime += result.elapsedTime;

                        BufferedImage newImg = new BufferedImage(img.getWidth(), img.getHeight(), img.getType());
                        newImg.setRGB(0, 0, img.getWidth(), img.getHeight(), result.pixels, 0, img.getWidth());
                        ImageIO.write(newImg, "jpg", new File("images/" + filename + "_proc_" + processor.getName() + ".jpg"));
                    }
                    averageTime /= TESTS_NUM;
                    System.out.printf("\t\t\tAverage time: %g s%n", averageTime);
                } catch (IOException e) {
                    System.out.printf("\tCould not read image. Details: %s%n", e.getMessage());
                }
            }
        }
    }
}
