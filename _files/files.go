package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// // we use os pacjage for operating system
	// // f,err:=os.Open("example.txt")
	// // if err!=nil{
	// // 	// we can log the error in here 
	// // 	/// we can put the program in panic here 
	// // 	panic(err)
	// // }
	// // f_info,err:=f.Stat()
	// // if err!=nil{
	// // 	// we can log the error in here 
	// // 	/// we can put the program in panic here 
	// // 	panic(err)
	// // }
	// // fmt.Println("file_name:",f_info.Sys())
	// // // isDir to check if it is a folder or file 
	// // //fileSize will give u the size of the file 
	// f,err:=os.Open("example.txt");

	// if err!=nil{
	// 	panic(err)
	// }

	// defer f.Close()
	// // now i wanna read it so i need to close it somewhere 
	// // buffer is an array of byte 
	// buf:=make([]byte,25);

	// d,err:=f.Read(buf);
	// if err!=nil{
	// 	panic(err)
	// }

	// for i:=0;i<len(buf);i++{
	// 	println("data",d,string(buf[i]));
	// }
	// println("data",buf,d)


	// // we can also use the ReadFile, but that's not good, as it loads the file at once. so we use 
	// //what if the fiel size is too long 

	// // when we gonna look upto the file size we can use straming 


	// now let's see how we gonna read, write and delete the file 

   // now let's read the file 
//    f,err:=os.Open("example.txt");

//    // now see if there is an error or not 
//    if err!=nil{
// 	 panic(err)
//    }
//    defer f.Close()   // why do we use it, so that at the last file is closed 
//    // after the main function ext 

//    // now i have to read the file opened 
//    // for that i would have to create buffer 

//    buf:=make([]byte,40)

//    d,err:=f.Read(buf)
  
// //    println("data",buf,d);

// //    for i,val=range buf{

// //    }
// // to read the data we woulf have to convert it tp string 
//   for i:=0;i<len(buf);i++{
// 	println("data",d,string(buf[i]));
//}

f,err:=os.ReadFile("example.txt");
if err!=nil{
	panic(err)
}
// f.Close(); we don;t need to opena nd close the file here 
// we don;t specifically use it 
// because the file maybe too large in here 
fmt.Println(string(f))


// we can aslo stream the file as we do in the case of nodejs 
//read folders 
  dir,e:=os.Open("../")
  if e!=nil{
	panic(e)
  }

  defer dir.Close();

  fmt.Println("these are the folder contents in the root folder")
  fileInfo,err:=dir.ReadDir(-1);
   for _,r:=range fileInfo{
        fmt.Println(r,r.IsDir())
  }
 // fmt.Println(fileInfo)

 // now let's create a file 
fi,er:=os.Create("example2.txt")
if er!=nil{
	panic(er)
}
// now see what can be done in here 
// let's hope for the best 
defer fi.Close();


// now lets do one thing read from one file and write to another file 
fii,ei:=os.Open("example.txt")
if ei!=nil{
	panic(ei)
}
defer fii.Close();

destFile,er:=os.Create("example3.txt");
if er!=nil{
	panic(er)
}
 defer destFile.Close();
// here we have opened a file and deleted a file 

//let's create a reader 
reader:=bufio.NewReader(fii)
// it wiill create a reader 
// we also need a writer 
writer:=bufio.NewWriter(destFile)

// now we can pipe , the data from reader to writer 
// we can use io package for that
// io.Copy(writer,reader) // it will copy the data from reader to writer 
// but we need to flush the data to the file 
// writer.Flush() // it will flush the data to the file

// dang, we can also copy in other way 
for{
   b,err:=reader.ReadByte();
   if err!=nil{
	if err.Error()=="EOF"{
		fmt.Println("file copied successfully")
	}else{
		panic(err)
	}
	break;
   }
   writer.WriteByte(b);
}
 writer.Flush();


 // now let's see how we can delete a file
// os.Remove("example3.txt") // it will delete the file
}

// what is bufio?
// bufio is a package in go which provides buffered io operations
// what is buffered io?
// buffered io is a technique which allows us to read and write data in chunks instead of reading and writing data byte by byte 
// it is more efficient than reading and writing data byte by byte
