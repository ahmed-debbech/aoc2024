package main;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;


public class puzzle1{

    public static class Towel{
        public String towelName;

        public Towel(String name){
            towelName = name;
        }
        @Override
        public String toString(){
            return towelName;
        }
    }
    
    public static class Desgin{
        public String designSet;

        @Override
        public String toString(){
            return designSet;
        }
        public Desgin(String h){
            designSet = h;
        }
    }

    public static void main(String[] args){
        System.out.println("hello aoc");

        List<Towel> bank = new ArrayList<>();
        List<Desgin> designs = new ArrayList<>();

        try {
            File myObj = new File("input");
            Scanner myReader = new Scanner(myObj);

            String data = myReader.nextLine();

            
            for (String s : data.split(",")){
                bank.add(new Towel(s));
            }
            
            myReader.nextLine();

            while (myReader.hasNextLine()) {
                String data1 = myReader.nextLine();
                designs.add(new Desgin(data1));
            }
            myReader.close();
        } catch (FileNotFoundException e) {
            System.out.println("An error occurred.");
            e.printStackTrace();
        }

        getCompatible(designs.get(0).designSet, bank);

    }

    public static List<Boolean> getCompatible(String design, List<Towel> bank){
        List<Boolean> comp = new ArrayList<>();

        for(int i=0; i<=design.length()-1; i++){
            comp.add(false);
        }


        for(Towel to : bank){
            for(int i=0; i<=comp.size()-1; i++){
                if(design.substring(i, to.towelName.length()).equals(to.towelName)){
                    for(int j = i ; j<= comp.size()-1; j++){
                        comp.set(j, true);
                    }
                }
            }
        }

        return comp;
    }
}