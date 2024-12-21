#include <iostream>
#include <vector>
#include <sstream>
#include <fstream>
#include <regex>

int main(){
    auto fpath = "./input.txt";
    auto file = std::fstream(fpath);
    
    std::cout << "parsing inputs " << std::endl;
    std::stringstream buffer;
    buffer << file.rdbuf();
    std::string text = buffer.str();

    std::basic_regex re("mul\\((\\d+),(\\d+)\\)");
    auto words_begin = std::sregex_iterator(
        text.begin(), text.end(), re
    );
    auto words_end = std::sregex_iterator();


    int result = 0;

    for (std::sregex_iterator i = words_begin; i != words_end; ++i) {
        std::smatch match = *i;
        result += std::stoi(match[1]) * std::stoi(match[2]);
    }

    std::cout << "Part 1: " << result << std::endl;

    // 
    std::basic_regex do_dont_re("(do\\(\\))|(don't\\(\\))|mul\\((\\d+),(\\d+)\\)");
    words_begin = std::sregex_iterator(
        text.begin(), text.end(), do_dont_re
    );
    words_end = std::sregex_iterator(); 
    bool enabled = true;
    int result_2 = 0;
    for (std::sregex_iterator i = words_begin; i!= words_end; ++i){
        std::smatch match = *i;
        // std::cout << "Found: " << match[0] << "\n";
        if (match[0] == "don't()") {
            enabled = false;
        }
        else if (match[0] == "do()"){
            enabled = true;
        }
        else if (!enabled){}
        else{
            result_2 += std::stoi(match[3]) * std::stoi(match[4]);  
        }

    }
    std::cout << "Part 2: " << result_2 << std::endl;
    return 0;
}