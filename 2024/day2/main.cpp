#include <iostream>
#include <vector>
#include <sstream>
#include <fstream>
#include <boost/numeric/ublas/vector.hpp>

typedef std::vector<int> Level;
typedef boost::numeric::ublas::vector<int> BoostVec;

bool is_small_diff(Level level, int tolerance){
    bool result = true;
    for (int i=0; i<level.size()-1;i++){
        int diff = level[i+1] -  level[i];
        if (std::abs(diff) > tolerance){
            result=false;
        };
    }
    return result;
}
bool is_increasing(Level level, int tolerance){
    int err_count = 0;
    for (int i=0; i<level.size()-1;i++){
        if (level[i+1] <= level[i]){
            err_count += 1;
        } 
    }
    return err_count <= tolerance;
}

bool is_decreasing(Level level, int tolerance){
    int err_count = 0;
    for (int i=0; i<level.size()-1;i++){
        if (level[i+1] >= level[i]){
            // std::cout << level[i+1] << level[i] << std::endl;
            err_count +=1;
        } 
    }
    return err_count <=tolerance;
}


bool validate_level(Level level, int tolerance){
    bool increasing = is_increasing(level, tolerance);
    bool decreasing = is_decreasing(level, tolerance);
    bool small_diff = is_small_diff(level, 3);
    // std::cout << "increasing: " << increasing << " decresing: " << decreasing << " small_diff: " << small_diff << std::endl;;
    return (increasing | decreasing) & small_diff  ;
}


bool validate_level_with_dampener(Level level) {
    // First check if it's valid without removing anything
    if (validate_level(level, 0)) {
        return true;
    }
    
    // Try removing each level one at a time
    for (int i = 0; i < level.size(); i++) {
        Level temp = level;
        temp.erase(temp.begin() + i);  // Remove ith element
        
        // Check if valid after removing this element
        if (validate_level(temp, 0)) {
            return true;
        }
    }
    return false;
}

std::vector<Level> parse_inputs(std::string fpath){
    auto input_file = std::fstream(fpath);
    std::cout << input_file.is_open() << std::endl;
    std::string line;
    
    std::cout << "parsing inpiuts " << std::endl;
    std::vector<Level> levels;
    while (std::getline(input_file, line)) {
    Level level;        
    std::istringstream ss(line);
    int val;
    while (ss >> val) {  // Only loops while extraction succeeds
            level.push_back(val);
    }
    levels.push_back(level);
}
    return levels;

}

void print_level(Level level){
    for (const auto& val: level){
        std::cout << val << ",";
    }
    std::cout << std::endl;
}
int main(){
    
    std::vector<Level> levels = parse_inputs("./input.txt");

    int tolerance=0;
    int n_valid = 0; 
    for (const auto& level : levels){
        if (validate_level(level, 0)){
            n_valid+=1;
        }
    }
    std::cout << "Part 1: " << n_valid << std::endl;
    

    n_valid = 0; 
    for (const auto& level : levels){
        // print_level(level);
        if (validate_level_with_dampener(level)){
            n_valid+=1;
        }
    }
    std::cout << "Part 2: " << n_valid << std::endl;
    return 0;
}
