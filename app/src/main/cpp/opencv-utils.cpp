#include "opencv-utils.h"
#include <opencv2/imgproc.hpp>
#include <thread>

void myBlur(Mat& image, float sigma) {
    Mat blurredImage;
    GaussianBlur(image, blurredImage, Size(0, 0), sigma);

    Mat bokehImage;
    subtract(image, blurredImage, bokehImage);

    threshold(bokehImage, bokehImage, 0, 255, THRESH_TOZERO);

    add(blurredImage, bokehImage, image);

}
void myNoise(Mat image, float sigma) {
    std::vector<Mat> channels;
    split(image, channels);

    for (Mat& channel : channels) {
        Mat noise(channel.size(), CV_32F);
        randn(noise, 0.0, sigma*3);

        Mat blurred;
        GaussianBlur(noise, blurred, Size(5, 5), 0);

        Mat result;
        add(channel, blurred, result, Mat(), CV_8U);

        channel = result;
    }
    merge(channels, image);
}

void myTone(Mat& image, float sigma)
{
    cvtColor(image, image, cv::COLOR_BGR2HSV);

    int hueShift = static_cast<int>((sigma / 10.0) * 180.0);


    for (int i = 0; i < image.rows; i++)
    {
        for (int j = 0; j < image.cols; j++)
        {
            auto& pixel = image.at<Vec3b>(i, j);
            int hue = static_cast<int>(pixel[0]) + hueShift;

            if (hue < 0)
                hue = 180 + hue;
            else if (hue > 179)
                hue = hue - 180;

            pixel[0] = static_cast<uchar>(hue);
        }
    }

    cvtColor(image, image, COLOR_HSV2BGR);
}

void mySaturation(Mat& image, float sigma) {
    cvtColor(image, image, COLOR_BGR2HSV);

    std::vector<Mat> channels;
    split(image, channels);

    float saturationShift = sigma / 10.0F;

    channels[1] *= (1.0 + saturationShift);

    merge(channels, image);

    cvtColor(image, image, COLOR_HSV2BGR);
}


void myBright(Mat image, float sigma) {
    Mat dst;
    float alpha = 1.0f, beta = sigma*3;

    image.convertTo(dst, -1, alpha,beta);
    convertScaleAbs(dst, image);
}

void myExposition(Mat& image, float sigma) {
    float exposition = sigma / 10.0F;

    Mat lookup_table(1, 256, CV_8U);

    for (int i = 0; i < 256; i++) {
        int adjusted_value = saturate_cast<uchar>(i + i * exposition);
        lookup_table.at<uchar>(i) = adjusted_value;
    }

    LUT(image, lookup_table, image);
}

void myContrast(Mat& image, float sigma) {
    float contrast = (sigma + 10.0F) / 10.0F;

    image.convertTo(image, -1, contrast, 0);
}

void myVignette(Mat& image, float sigma) {
    cvtColor(image, image, COLOR_BGR2HSV);

    int width = image.cols;
    int height = image.rows;

    int centerX = width / 2;
    int centerY = height / 2;

    double maxDistance = norm(Point(centerX, centerY));

    for (int y = 0; y < height; ++y)
    {
        for (int x = 0; x < width; ++x)
        {
            double distance = norm(cv::Point(x - centerX, y - centerY));
            double vignette = 1.0 - sigma * distance / maxDistance;

            vignette = std::max(vignette, 0.0);

            image.at<Vec3b>(y, x)[2] *= vignette;

            image.at<Vec3b>(y, x)[1] *= vignette;
        }
    }

    cvtColor(image, image, COLOR_HSV2BGR);
}
void myAutocorrect(Mat& image) {
    Mat floatImage;
    image.convertTo(floatImage, CV_32F, 1.0 / 255.0);

    std::vector<Mat> channels;
    split(floatImage, channels);

    Scalar meanR = mean(channels[2]);
    Scalar meanG = mean(channels[1]);
    Scalar meanB = mean(channels[0]);

    float scaleR = meanG[0] / meanR[0];
    float scaleB = meanG[0] / meanB[0];

    channels[2] = channels[2] * scaleR;
    channels[0] = channels[0] * scaleB;
    channels[1] = channels[1] * 1.0;

    merge(channels, floatImage);

    floatImage.convertTo(image, CV_8U, 255.0);
}



