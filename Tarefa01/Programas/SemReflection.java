package com.mycompany.reflection;

import java.lang.reflect.Field;

public class Reflection {
    
   static class Pessoa {

        private double peso;
        private double altura;

        public Pessoa(double peso, double altura){
            this.peso=peso;
            this.altura=altura;

        }

        public double getPeso(){
            return this.peso;

        }

        public double getAltura(){
            return this.altura;

        }
        
        public void setAltura(double altura){
            this.altura=altura;
        }
        
        public void setPeso(double peso){
            this.peso=peso;
        }        

        public double imc(){
            return peso/(altura*altura);

        }

    } 

    public static void main(String[] args) {
        Pessoa pessoa1 = new Pessoa(90,1.9);
        System.out.println("O IMC é: " + pessoa1.imc());
        pessoa1.setPeso(80);
        pessoa1.setAltura(1.8);
        System.out.println("O novo IMC é: " + pessoa1.imc());
    
    }
    




}
