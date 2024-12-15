#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <unordered_map>

int compute_l1(std::vector<int> x, std::vector<int> y){
    int sum_l1_diff = 0;
    for (int i=0;i<x.size();i++){
        int diff = std::abs(x[i] - y[i]);
        sum_l1_diff += diff;
    }
    return sum_l1_diff;
}

typedef std::unordered_map<int, int> Counter;

Counter create_counter(std::vector<int> x)
{
    Counter counter = {};
    for (const auto& item : x) {
        counter[item]++;
    }
    return counter;
}

int compute_similarity(Counter x, Counter y){
    int result = 0;
    for (const auto& [key, freq]: x){
        // std::cout << "num:" << key << "freq:" << freq << "," << y[key] << std::endl;
        result += key * freq * y[key];
    }
    return result;
}

int main(){
    std::cout << "hello world" << std::endl;
    auto input_file = std::fstream("./input.txt");
    // std::cout << input_file.is_open() << std::endl;
    std::string line;
    

    std::vector<int> lhs, rhs; 

    while (std::getline(input_file, line)) {
    std::istringstream ss(line);
    int first, second;
    ss >> first >> second;
    lhs.push_back(first);
    rhs.push_back(second);
    }

    std::sort(lhs.begin(), lhs.end());
    std::sort(rhs.begin(), rhs.end());

    // for (const auto& item: lhs){
        // std::cout << item << std::endl;
    // }

    int result = compute_l1(lhs, rhs); 
    std::cout << "Part 1 Result: " << result << std::endl;

    auto counter_x = create_counter(lhs);
    auto counter_y = create_counter(rhs);
    auto similarity = compute_similarity(counter_x, counter_y);
    std::cout << "Part 2 Result: " << similarity << std::endl; 
}   

