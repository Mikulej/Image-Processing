#include <iostream>
#include <fstream>
#include <list>

struct Pixel
{
	int R, G, B, A;
};

struct Bitmap
{
	int width, height;
	std::list<Pixel> pixels;
};

Bitmap readPNG(std::string inputFilename)
{
	std::ifstream inStream(inputFilename, std::ios::out | std::ios::binary);
	if (!inStream)
	{
		std::cout << "Error while opening" << std::endl;
		return {};
	}

	char inHeader[8];
	inStream.read(inHeader, 8);
	char expectedHeader[8] = {0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A};
	for (int i = 0; i < 8; ++i)
	{
		if (inHeader[i] != expectedHeader[i])
		{
			inStream.close();
			std::cout << "Error while reading file (invalid PNG header)" << std::endl;
			return {};
		}
	}

	int width = 0, height = 0;
	bool AdamInterlace = false;

	int chunksLoaded = 0;
	//while (!inStream.eof()) {
	/*char testBytes[29];
	inStream.read(testBytes, 29);
	for (int i = 0; i < 29; ++i) {
		std::cout << "[" << i << "] Dec: " << (int)(testBytes[i]);
		if (isalnum((unsigned char)(testBytes[i]))) {
			std::cout << ", Char: " << (unsigned char)(testBytes[i]);
		}
		std::cout << std::endl;
	}*/
	char chunkLengthBytes[4];
	char chunkTypeBytes[5] = {0};
	inStream.read(chunkLengthBytes, 4);
	inStream.read(chunkTypeBytes, 4);
	int chunkLength = chunkLengthBytes[0] * 256 * 256 * 256 + chunkLengthBytes[0] * 256 * 256 + chunkLengthBytes[2] *
		256 + chunkLengthBytes[3];

	if (chunksLoaded == 0)
	{
		if (chunkLength != 13)
		{
			inStream.close();
			std::cout << "Error while reading file (declared length of first chunk is " << chunkLength <<
				" bytes instead of 13)" << std::endl;
			return {};
		}

		for (int i = 0; i < 4; ++i)
		{
			if (chunkTypeBytes[i] != "IHDR"[i])
			{
				inStream.close();
				std::cout << "Error while reading file (first chunk is named " << chunkTypeBytes << " instead of IHDR)"
					<< std::endl;
				return {};
			}
		}
	}

	unsigned char widthBytes[4], heightBytes[4], paramBytes[5];

	inStream.read((char*)widthBytes, 4);
	width = widthBytes[0] * 256 * 256 * 256 + widthBytes[1] * 256 * 256 + widthBytes[2] * 256 + widthBytes[3];
	inStream.read((char*)heightBytes, 4);
	height = heightBytes[0] * 256 * 256 * 256 + heightBytes[1] * 256 * 256 + heightBytes[2] * 256 + heightBytes[3];
	inStream.read((char*)paramBytes, 5);

	int const bitDepth = paramBytes[0];
	int const colorType = paramBytes[1];
	if (paramBytes[2] != 0)
	{
		std::cout << "Warning: Compression method byte is not equal to 0" << std::endl;
	}
	if (paramBytes[3] != 0)
	{
		std::cout << "Warning: Filtering method byte is not equal to 0" << std::endl;
	}
	AdamInterlace = paramBytes[4];

	std::cout << "Loaded image parameters:" << std::endl;
	std::cout << "Dimensions: " << width << "x" << height << std::endl;
	std::cout << "Bit depth per sample or palette index: " << bitDepth << std::endl;
	std::string colorTypeInfos[7] = {
		"Grayscale", //0
		"<invalid type>", //1
		"RGB values", //2
		"Palette index values", //3
		"Grayscale + alpha", //4
		"<invalid type>", //5
		"RGB + Alpha", //6
	};
	std::string colorTypeString = ((colorType >= 0 && colorType <= 6) ? colorTypeInfos[colorType] : "<invalid type>");
	std::cout << "Color type: " << colorType << " (" << colorTypeString.c_str() << ")" << std::endl;

	std::cout << "Interlace: " << (AdamInterlace ? "Adam7" : "none") << std::endl;

	chunksLoaded++;
	//}

	inStream.close();
	if (!inStream.good())
	{
		std::cout << "Opened but not read" << std::endl;
		return {};
	}
	return {width, height, {{1, 2, 3, 4}, {5, 6, 7, 8}, {2, 2, 3, 3}}};
}

int main()
{
	/*std::ofstream outStream("out_filename.bmp", std::ios::out | std::ios::binary);
	if (!outStream) {
		std::cout << "Error while opening file" << std::endl;
		getchar();
		return 1;
	}
	char bytes[4] = { 'a', 'b', 'c', 'd' };
	std::cout << "Probujemy zapisac bajty:" << std::endl;
	for (int i = 0; i < 4; ++i) {
		std::cout << (int)(bytes[i]) << std::endl;
	}
	outStream.write(bytes, 4);
	outStream.close();
	if (!outStream.good()) {
		std::cout << "Opened but not saved" << std::endl;
		getchar();
		return 1;
	}*/

	Bitmap loadedBitmap = readPNG("in_filename.png");
	if (loadedBitmap.width != 0)
	{
		std::cout << "File loaded correctly" << std::endl;
	}

	getchar();
	return 0;
}
