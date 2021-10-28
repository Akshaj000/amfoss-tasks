import java.util.ArrayList;
import java.util.Scanner;

public class main {

    public static void print(int hash, int previoushash , String[] message ,int  n){
        System.out.println("block-"+n);
        System.out.println("Hash:");
        System.out.println(hash);
        System.out.println("Previous Hash:");
        System.out.println(previoushash);
        System.out.println("Message:");
        System.out.println(message[0]);
        System.out.println(" ");
    }

    ArrayList<block> blockchain = new ArrayList<>();
    public static <string> void main(String[] args) {
        System.out.println("Enter no of blocks:");
        Scanner input = new Scanner(System.in);
        int t = input.nextInt();
        String[] messages = new String[t];

        for (int i=0 ; i<t ; i++){
            System.out.println("Enter message in "+i+"th block:");
            Scanner block = new Scanner(System.in);
            messages[i]=block.nextLine();
        }

        while (true){
            String[] genesisTransactions = {messages[0]};
            block genesisblock = new block( 0, genesisTransactions);
            print(genesisblock.getBlockHash(), 0,genesisTransactions,0);
            int  pblck = genesisblock.getBlockHash();
            String[] ptrans = genesisTransactions;
            for (int i=1 ; i<t ; i++){
                String[] Transactions = {messages[i]};
                block nblock = new block( pblck, Transactions);
                print(nblock.getBlockHash(), pblck,Transactions,i);
                pblck = nblock.getBlockHash();
                ptrans = Transactions;
            }
            System.out.println("Enter no of block to change the message or -1 to exit: ");
            Scanner inp = new Scanner(System.in);
            int option = inp.nextInt();
            if(option == -1){
                System.exit(0);
            }
            else{
                if (option >= 0 && option < t){
                    System.out.println("Enter new message:");
                    Scanner message = new Scanner(System.in);
                    String mess = message.nextLine();
                    messages[option] = mess;
                }
            }
        }

    }
}
