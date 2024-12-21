#include <iostream>
#include <vector>
#include <sstream>
#include <fstream>
#include <regex>

int main(){
    std::cout << "hello world" << std::endl;
    auto fpath = "./test_input.txt";
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
        // match[0] is the complete match
        // match[1] is the first number
        // match[2] is the second number
        std::cout << "Found: " << match[0] << "\n";
        std::cout << "  First number: " << match[1] << "\n";
        std::cout << "  Second number: " << match[2] << "\n";
        result += std::stoi(match[1]) * std::stoi(match[2]);
    }

    std::cout << result << std::endl;

    std::basic_regex do_dont_re("do\\(\\)|(don't\\(\\))|mul\\((\\d+),(\\d+)\\)");
    words_begin = std::sregex_iterator(
        text.begin(), text.end(), do_dont_re
    );
    words_end = std::sregex_iterator(); 
    bool enabled = true;
    result = 0;
    for (std::sregex_iterator i = words_begin; i!= words_end; ++i){
        std::smatch match = *i;
        std::cout << "Found: " << match[0] << "\n";
        if (match[0] == "don't()") {
            enabled = false;
            std::cout << "Found dont" << std::endl;
        }
        else if (match[0] == "do()"){
            enabled = true;
            std::cout << "Found do" << std::endl;
        }
        else if (!enabled){}
        else{
            std::cout << "  First number: " << match[3] << "\n";
            std::cout << "  Second number: " << match[4] << "\n";
            result += std::stoi(match[3]) * std::stoi(match[4]);  
        }

    }
    std::cout << "part 2: " << result << std::endl;
    return 0;
}