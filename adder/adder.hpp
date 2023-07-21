#ifndef _ADDER_H_
#define _ADDER_H_
class Adder{
  private:
    int count = 0;    
    
  public:
    void Add(int num){
      count += num;
    };
        
    int Total(){
      return count;
    };


};
#endif
