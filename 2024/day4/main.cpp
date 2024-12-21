#include <iostream>
#include <fstream>
#include <vector>
#include <tuple>

typedef std::vector<std::vector<char>> Grid;
typedef std::tuple<int, int> vec;

std::vector<vec> directions = {{1,0}, {0,1}, {-1,0}, {0,-1}, {1,1}, {1,-1}, {-1,1}, {-1,-1}};


class Board {
    std::vector<std::vector<char>> grid;


    public:
    Board(std::vector<std::vector<char>> grid_){
        grid = grid_;
    }

    void print(){
        for (const auto& row: grid){
            for (const auto& item: row){
                std::cout << item;
            }
            std::cout << std::endl;
        }
    }

    int part_one(std::string word){
        int count = 0;
        for (int i = 0; i<grid.size(); i++){
            for (int j = 0; j<grid[i].size(); j++){
                // std::cout << "starting with i=" << i << " j=" << j <<std::endl; 
                count += count_matches(i, j, word);
            }
        }
        return count;
    }

    int count_matches(int i,int j, std::string word){
        int count = 0;
        for (const auto& direction : directions){
            if (attempt_directional_match(i,j,word,direction)){
                count +=1;
            }
        }
        return count;
    }

    bool validate_point(int i,int j){
        if (i <0||i>=grid.size()){
            return false;
        }
        if (j<0 || j>=grid[0].size()){
            return false;
        } 
        return true;
    }

    bool attempt_directional_match(int i, int j, std::string word, std::tuple<int, int> direction){
            auto [di, dj] = direction;
            for (int l =0; l<word.size(); l++) {
                int i_candidate = i +l*di;
                int j_candidate = j + l*dj;
                if (!validate_point(i_candidate, j_candidate)){
                    return false;
                }
                char candidate  = grid[i_candidate][j_candidate];
                if ( candidate != word.at(l)) {
                    return false;
                }
            }
        return true;
    }

    int part_two(){
        int count = 0;
        for (int i = 0; i<grid.size(); i++){
            for (int j = 0; j<grid[i].size(); j++){
                if (is_x_mas(i, j)){
                    count++;
                }
            }
        }
        return count;
        
    }

    bool is_x_mas(int i, int j){
        if (grid[i][j] != 'A') {
            return false;
        }
        
        if (!validate_point(i+1,j+1)){
            return false;
        }
        if (!validate_point(i-1, j-1)){
            return false;
        }
        char a = grid[i+1][j-1];
        char b = grid[i+1][j+1];
        char c = grid[i-1][j+1];
        char d =  grid[i-1][j-1] ;
        if(((a == 'M' && c == 'S') || (a == 'S' && c == 'M')) && ((b == 'M' && d == 'S') || (b == 'S' && d == 'M'))){
            return true;
        }
        return false;
    }
};
    
int main(){
    auto fpath = "./input.txt";
    auto input_file = std::fstream(fpath);
    Grid raw_grid;
    std::string line;
    while (std::getline(input_file, line)) {
        std::vector<char> row(line.begin(), line.end());
        raw_grid.push_back(row);
    }

    Board board(raw_grid);
    // board.print();
    int part_one = board.part_one("XMAS");
    std::cout << "Part 1: " << part_one << std::endl;
    int part_two = board.part_two();
    std::cout << "Part 2: " << part_two << std::endl;
};