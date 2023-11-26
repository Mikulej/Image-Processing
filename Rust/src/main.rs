use image::DynamicImage;
use image::io::Reader as ImageReader;
use image::GenericImageView;
use std::fs;
use std::time::{SystemTime};


fn main() {
    let start =SystemTime::now();
    //let current_dir = env::current_dir().expect("Failed to get current directory");
    //println!("current_dir: {:?}", current_dir);
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
    //println!("{:?}", image_paths);

    //2.Load Image
    let mut img = load_image(&image_paths[0]);
    
    // Obtain the image's width and height.
    let (width, height) = img.dimensions();
    let mut imgbuf = image::ImageBuffer::new(width, height);
    //3. Manipulate image
    let mode = args[1].as_str();
        match mode {
            "r" => {}, //removal
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
                img = img.blur(blurStrength);
                img.save("out/out.png").unwrap();         
            }, 
            _ => panic!("Invalid operation. Expected 'r', 'm', or 'b'."),
        }
    println!("Time elapsed: {:?}", start.elapsed());
}

fn load_image(path: &str) -> DynamicImage {
    let reader = ImageReader::open(path).expect("Failed to open image file");
    reader.decode().expect("Failed to decode image")
}

