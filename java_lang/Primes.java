package java_lang;

// If a number is prime, number i.e, if a number is positive number and has 1 and itself only 
// as the positive divisors
// i.e, -1 -> false
// 0 -> false   6 -> false
// 1 -> false   7 -> true
// 2 -> true    8 -> false
// 3 -> true	9 -> false
// 4 -> false	10 -> false
// 5 -> true


public class Primes{
    public Boolean IsPrime(int num){
        if (num <= 1){
            return false;
        }

        int i = 2;
        
        while (i < num){
            if (num % i == 0) {
                return false;
            }
            i++;
        }
    
        return true;
    }

    public static void main(String[] args) {
        Primes primes = new Primes();
        for (int i = -1; i <= 10; i++){
            System.out.println(primes.IsPrime(i));
        }
    }
}
