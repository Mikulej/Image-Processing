
use image::DynamicImage;
use image::io::Reader as ImageReader;
use image::GenericImageView;
use std::env;
use std::fs;

fn main() {
    println!("Program started");
    let current_dir = env::current_dir().expect("Failed to get current directory");
    println!("current_dir: {:?}", current_dir);
    let args: Vec<String> = std::env::args().collect();
    println!("Argument count: {}", args.len());
    if args.len() == 2 {
        println!("Passed 1 argument");
    }
    println!("{:?}", args);
    //arg[0] path to image1
    //arg[1] path to image2
    //arg[2] operation r, m, b
    
    //0. Get image path
    let paths = fs::read_dir("./images").unwrap();

    let mut image_paths = Vec::new();

    for path in paths {
        let path_str = path.unwrap().path().display().to_string();
        if path_str.ends_with(".png") || path_str.ends_with(".jpg") {
            image_paths.push(path_str);
        }
    }

    println!("{:?}", image_paths);
    //1.Load Image
    let mut img = load_image(&image_paths[0]);
    // Obtain the image's width and height.
    let (width, height) = img.dimensions();
    //2.Convert to fixed data type
    //3. Manipulate image
    let mode = args[1].as_str();
        match mode {
            "r" => {}, //removal
            "m" => {
                let img2 = load_image(&image_paths[1]);
                img2.resize(width, height, image::imageops::FilterType::Nearest);                
            }, //merge
            "b" => {
                img = img.blur(10.0);

            }, //blur
            _ => panic!("Invalid operation. Expected 'r', 'm', or 'b'."),
        }
    //4. Save image
    save_image(&img, "out/out.png");
}

fn load_image(path: &str) -> DynamicImage {
    let reader = ImageReader::open(path).expect("Failed to open image file");
    reader.decode().expect("Failed to decode image")
}

fn save_image(image: &DynamicImage, path: &str) {
    image
        .save(path)
        .expect("Failed to save image");
}
