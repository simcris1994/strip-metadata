# strip-metadata

This is an API which allows you to upload an image and responds with the same image, stripped of EXIF data.

# How to run

1. You need to have Go installed on your machine
2. Download this repo
3. Run "go run ."
4. Use curl or a REST client to call POST http://localhost:8080/upload. Set the Content-Type header to multipart/form-data and provide the file you want to have processed in the body of the request. The key of the entry should have type File and be named "file".
5. Save the image you get in the response and you're done. 
