//Assignment 1
//Course Code : Csc 220
//Course Title : Data Structure
//Department : Mathermatics

void main() {
  List<String> studentNames = [
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

  List<int> studentId = [
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

  List<double> studentGp = [
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

  // print(studentGp.length);

  //this method call displays the size of the array
  getArraySize(studentNames);

//this method call displays the all the elements in the array
  getArrayList(studentNames);

  //this method call searches for an element in the array
  // searchElement(studentNames, "jeebolar");

  //this method call removes an element in the array
  // removeElement(studentNames, "maryam");

  // getArrayList(studentNames);
}

//this method gets the length of the list passed in as parameter
void getArraySize(List array) {
  print(array.length);
}

//this method loops or cycles through the list
//and prints all the elements in the list

void getArrayList(List array) {
  print("List of elements in the array");
  for (int i = 0; i < array.length; i++) {
    print(array[i]);
  }
}

//method to search for an item

void searchElement(List array, var item) {
  for (int i = 0; i < array.length; i++) {
    if (array[i] == item) {
      print(item + " found in index $i of the array");
    }
  }
}

//this methods removes an element from the list
void removeElement(List array, var item) {
  for (int i = 0; i < array.length; i++) {
    if (array[i] == item) {
      array.removeAt(i);
      print(item + " removed from the student Names");
    }
  }
}
