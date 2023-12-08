use image::DynamicImage;
use image::io::Reader as ImageReader;
use image::GenericImageView;
use std::fs;


fn main() {
   
    let args: Vec<String> = std::env::args().collect();
    
    //1. Get images' paths
    let paths = fs::read_dir("./images").unwrap();

    let mut image_paths = Vec::new();

    for path in paths {
        let path_str = path.unwrap().path().display().to_string();
        if path_str.ends_with(".jpg") {
            image_paths.push(path_str);
        }
    }
   
        //2.Load Image
        let img = load_image(&image_paths[0]);
        let (width, height) = img.dimensions(); // Obtain the image's width and height.
        let mut imgbuf = image::ImageBuffer::new(width, height);
        //3. Manipulate image
    
        let mode = args[1].as_str();
            match mode {
                "r" => {//removal
                    imgbuf = img.to_rgba8();           
                    for i in (2..args.len()).step_by(4) {
                        if let Some(arg1) = args.get(i) {
                            if let Some(arg2) = args.get(i + 1) {
                                if let Some(arg3) = args.get(i + 2) {
                                    if let Some(arg4) = args.get(i + 3){
                                        match arg1.as_str(){
                                            "s" =>{//square
                                                let x = arg2.parse::<i32>().unwrap();
                                                let y = arg3.parse::<i32>().unwrap();
                                                let length = arg4.parse::<u32>().unwrap();
                                                for i in x..(x+length as i32){
                                                    if i >= width as i32 || i < 0{continue;}
                                                    for j in y..(y+length as i32){
                                                        if j >= height as i32 || j < 0 {continue;}
                                                        imgbuf.put_pixel(i as u32, j as u32, image::Rgba([0, 0, 0, 0]));
                                                    }
                                                }
                                            },
                                            "c" =>{//circle
                                                let x = arg2.parse::<i32>().unwrap();
                                                let y = arg3.parse::<i32>().unwrap();
                                                let radius = arg4.parse::<u32>().unwrap();
                                                for i in (x-radius as i32)..(x+radius as i32){
                                                    if i >= width as i32 || i < 0 {continue;}
                                                    for j in (y-radius as i32)..(y+radius as i32){
                                                        if j >= height as i32 || j < 0 {continue;}
                                                        if (i-x)*(i-x) + (j-y)*(j-y) <= (radius * radius) as i32{
                                                            imgbuf.put_pixel(i as u32, j as u32, image::Rgba([0, 0, 0, 0]));
                                                        }
                                                    }
                                                }
                                            },
                                            _=> panic!("Invalid shape. Expected 's' or 'c'. Was {}", arg1),
                                        }
                                    }          
                                } 
                            } 
                        } 
                    }
                
                    imgbuf.save("out/out.png").unwrap();
                }, 
                "m" => { //merge
                    let ratio1 = args[2].parse::<f32>().unwrap();
                    let ratio2 = 1.0 - ratio1;
                    let mut img2 = load_image(&image_paths[1]);
                    let (width2, height2) = img2.dimensions();                    
                    for(x, y, pixel1) in img.pixels() {
                        if x >= width2 || y >= height2 {continue;} //if the pixel is out of bounds, skip it
                        let mut pixel2 = img2.get_pixel(x, y);
                        let mut mergedPixel: image::Rgba<u8> = image::Rgba([
                            ((pixel1[0] as f32 * ratio1) + (pixel2[0] as f32 * ratio2)) as u8,
                            ((pixel1[1] as f32 * ratio1) + (pixel2[1] as f32 * ratio2)) as u8,
                            ((pixel1[2] as f32 * ratio1) + (pixel2[2] as f32 * ratio2)) as u8,
                            ((pixel1[3] as f32 * ratio1) + (pixel2[3] as f32 * ratio2)) as u8
                            ]);
                        imgbuf.put_pixel(x, y, mergedPixel);
                    } 
                    imgbuf.save("out/out.png").unwrap();
                }, 
                "b" => {//blur
                    for(x, y,pixel) in img.pixels(){
                        if x <= 1 || y <= 1 || x >= width-2 || y >= height-2{continue;}
                        let nearbyPixels = [
                            img.get_pixel(x-2, y-2),
                            img.get_pixel(x-2, y-1),
                            img.get_pixel(x-2, y),
                            img.get_pixel(x-2, y+1),
                            img.get_pixel(x-2, y+2),
                            img.get_pixel(x-1, y-2),
                            img.get_pixel(x-1, y-1),
                            img.get_pixel(x-1, y),
                            img.get_pixel(x-1, y+1),
                            img.get_pixel(x-1, y+2),
                            img.get_pixel(x, y-2),
                            img.get_pixel(x, y-1),
                            pixel,
                            img.get_pixel(x, y+1),
                            img.get_pixel(x, y+2),
                            img.get_pixel(x+1, y-2),
                            img.get_pixel(x+1, y-1),
                            img.get_pixel(x+1, y),
                            img.get_pixel(x+1, y+1),
                            img.get_pixel(x+1, y+2),
                            img.get_pixel(x+2, y-2),
                            img.get_pixel(x+2, y-1),
                            img.get_pixel(x+2, y),
                            img.get_pixel(x+2, y+1),
                            img.get_pixel(x+2, y+2)
                        ];
                        let mut blurredPixel: image::Rgba<u8> = pixel;

                        for i in 0..3{
                            blurredPixel[i] = (
                                nearbyPixels[0][i] as f32 * 0.00390625 + //1
                                nearbyPixels[1][i] as f32 * 0.015625 +   //4
                                nearbyPixels[2][i] as f32 * 0.0234375 +   //6
                                nearbyPixels[3][i] as f32 * 0.015625 +   //4  
                                nearbyPixels[4][i] as f32 * 0.00390625 +//1
                                nearbyPixels[5][i] as f32 * 0.015625 +//4
                                nearbyPixels[6][i] as f32 * 0.0625 +//16
                                nearbyPixels[7][i] as f32 * 0.09375 +//24
                                nearbyPixels[8][i] as f32 * 0.0625 +//16
                                nearbyPixels[9][i] as f32 * 0.015625 +//4
                                nearbyPixels[10][i] as f32 * 0.0234375 +//6
                                nearbyPixels[11][i] as f32 * 0.09375 +//24
                                nearbyPixels[12][i] as f32 * 0.140625 +//36
                                nearbyPixels[13][i] as f32 * 0.09375 +//24
                                nearbyPixels[14][i] as f32 * 0.0234375 +//6
                                nearbyPixels[15][i] as f32 * 0.015625 +//4
                                nearbyPixels[16][i] as f32 * 0.0625 +//16
                                nearbyPixels[17][i] as f32 * 0.09375 +//24
                                nearbyPixels[18][i] as f32 * 0.0625 +//16
                                nearbyPixels[19][i] as f32 * 0.015625 +//4
                                nearbyPixels[20][i] as f32 * 0.00390625 +//1
                                nearbyPixels[21][i] as f32 * 0.015625 +//4
                                nearbyPixels[22][i] as f32 * 0.0234375 +//6
                                nearbyPixels[23][i] as f32 * 0.015625 +//4
                                nearbyPixels[24][i] as f32 * 0.00390625//1
                            ) as u8;
                        }
                        imgbuf.put_pixel(x, y, blurredPixel);
                    }   
              
                    imgbuf.save("out/out.png").unwrap();      
                    
                }, 
                _ => panic!("Invalid operation. Expected 'r', 'm', or 'b'. Was {}", mode),
            }
    //TESTS
    // let mut total_time = std::time::Duration::new(0, 0);
    // for _ in 0..12{
    // let start = std::time::Instant::now();
    //     let end = std::time::Instant::now();
    //     let elapsed_time = end.duration_since(start);
    //     total_time += elapsed_time;
    // }
    // println!("Total time elapsed: {:?}", total_time);
    
}

fn load_image(path: &str) -> DynamicImage {
    let reader = ImageReader::open(path).expect("Failed to open image file");
    reader.decode().expect("Failed to decode image")
}

