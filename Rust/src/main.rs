use image::DynamicImage;
use image::io::Reader as ImageReader;
use image::GenericImageView;
use std::fs;
use std::time::{SystemTime};


fn main() {
    let start =SystemTime::now();
    let args: Vec<String> = std::env::args().collect();
    
    //1. Get images' paths
    let paths = fs::read_dir("./images").unwrap();

    let mut image_paths = Vec::new();

    for path in paths {
        let path_str = path.unwrap().path().display().to_string();
        if path_str.ends_with(".png") || path_str.ends_with(".jpg") {
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
                                                if(i >= width as i32 || i < 0){continue;}
                                                for j in y..(y+length as i32){
                                                    if(j >= height as i32 || j < 0){continue;}
                                                    imgbuf.put_pixel(i as u32, j as u32, image::Rgba([0, 0, 0, 0]));
                                                }
                                            }
                                        },
                                        "c" =>{//circle
                                            let x = arg2.parse::<i32>().unwrap();
                                            let y = arg3.parse::<i32>().unwrap();
                                            let radius = arg4.parse::<u32>().unwrap();
                                            for i in (x-radius as i32)..(x+radius as i32){
                                                if(i >= width as i32 || i < 0){continue;}
                                                for j in (y-radius as i32)..(y+radius as i32){
                                                    if(j >= height as i32 || j < 0){continue;}
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
                img2 = img2.resize(width, height, image::imageops::FilterType::Nearest);  
               
                for(x, y, pixel1) in img.pixels() {
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
                let blurStrength = args[2].parse::<f32>().unwrap();
                imgbuf = img.blur(blurStrength).to_rgba8();
                imgbuf.save("out/out.png").unwrap();         
            }, 
            _ => panic!("Invalid operation. Expected 'r', 'm', or 'b'. Was {}", mode),
        }
    println!("Time elapsed: {:?}", start.elapsed());
}

fn load_image(path: &str) -> DynamicImage {
    let reader = ImageReader::open(path).expect("Failed to open image file");
    reader.decode().expect("Failed to decode image")
}

