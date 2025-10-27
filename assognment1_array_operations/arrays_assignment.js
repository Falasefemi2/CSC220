var  studentNames = [
    "ayo",
    "bola",
    "rufai",
    "jeebolar",
    "taofeeq",
    "muhammed",
    "boye",
    "ajayi",
    "maryam",
    "hassan"
  ];

 var studentId = [
    183749,
    373840,
    378493,
    383947,
    474959,
    374850,
    463872,
    485767,
    377384,
    475859,
  ];

  var studentGp = [
    3.56,
    4.56,
    4.32,
    5.00,
    2.45,
    3.56,
    2.45,
    2.56,
    1.34,
    4.34
  ];




//this displays all the element in the array
for (i = 0; i < studentId.length; i++) {
    console.log(studentId[i]);
}


//this adds an element to the array(studentId) with the push method and displays the array 
var newStudentId = 111111;
studentId.push(newStudentId);
for (i = 0; i < studentId.length; i++) {
    console.log(studentId[i]);
}

//this gets the size orlength of the array
console.log(studentId.length);


//this searches for an element in the array
for (i = 0; i < studentId.length; i++) {
    if(studentNames[i]=="maryam"){
        console.log( "maryam found in index  "+ i +"  in the studentNames")
    }
}


// this algorithm removes an element from the list
delete studentNames[2];
console.log(studentNames);



