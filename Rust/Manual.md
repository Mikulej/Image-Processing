## Beforing running
Create a folder named **images**  where Rust.exe is located and put two pictures there.   
In the same place create folder named **out** where program's output will appear.  
## How to use Rust.exe
First argument specifies mode:
- r = removal
- m = merge
- b = blur
### Removal
First argument specifies type of shape:
- c = Circle
- s = Square

Another two specify x and y of top-left corner of a square or center of a circle.  
The last argument narrows length of a side for square or radius of a circle.  
It's possible to use multiple shapes at once.  
Example:   
`rust.exe r c 20 50 100 s 350 400 150`
### Merge
Argument with range of [0;1], specifies percentage ratio of the first image, ratio of a second image is [1 - ratio of the first picture]  
Example:   
`rust.exe m 0.4`
### Blur
Argument with range of [0;INT_MAX], specifies strength of blur.  
Example:  
`rust.exe b 10`