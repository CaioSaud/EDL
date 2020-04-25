import java.lang.reflect.Field;

public class Reflection {
    
   static class Pessoa {

        private double peso;
        private double altura;

        public Pessoa(double peso, double altura){
            this.peso=peso;
            this.altura=altura;

        }


        public double imc(){
            return peso/(altura*altura);

        }

    } 

    public static void main(String[] args) {
        Pessoa pessoa1 = new Pessoa(90,1.9);
        System.out.println("O IMC é: " + pessoa1.imc());
        try{
            
            Field var1 = Pessoa.class.getDeclaredField("altura");
            Field var2 = Pessoa.class.getDeclaredField("peso");
            var1.setAccessible(true);
            var2.setAccessible(true);
            var1.setDouble(pessoa1,1.9);
            var2.setDouble(pessoa1,90);
            var1.setAccessible(false);
            var2.setAccessible(false);
            
        }catch(Exception err){
            err.printStackTrace();
        }
        
        System.out.println("O novo IMC é: " + pessoa1.imc());
    
    }
    




}
