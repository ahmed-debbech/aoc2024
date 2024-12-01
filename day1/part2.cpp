#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <algorithm>
#include <cmath>

using namespace std;

int main(){
	cout << "Hello AOC" << endl;
	ifstream file("input");
	string line;
	vector<int> left_list;
	vector<int> right_list;
	string r = "";
	string l = "";
	while(getline(file, line)){
		//cout << line << endl;
		l += line[0]; l+= line[1]; l+= line[2]; l+= line[3]; l+= line[4];
		r += line[8]; r+= line[9]; r+= line[10]; r+= line[11]; r+= line[12];

		//cout << "l " << l << " r " << r << endl; 	
		left_list.push_back(atoi(l.c_str()));
		right_list.push_back(atoi(r.c_str()));
		r = ""; l ="";
	}
	file.close();
	
	long int result = 0;
	for(int i=0; i<=left_list.size()-1; i++){
		int occ = 0;
		for(int j=0; j<=right_list.size()-1; j++){
			if(left_list[i] == right_list[j]){
				occ++;
			}
		}
		result += (left_list[i] * occ);
	}
	cout << "THE RESULT FOR PART 2 IS: " << result << endl;
	return 0;
}
