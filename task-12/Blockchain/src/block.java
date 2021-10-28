import java.util.Arrays;

public class block {

    private int previousHash;
    private String[] transaction;

    private  int blockHash;
    private String[] transactions;

    public block(int previousHash, String[] transactions) {
        this.previousHash = previousHash;
        this.transactions = transactions;

        Object[] contents  = {Arrays.hashCode(transactions),previousHash};
        this.blockHash = Arrays.hashCode(contents);
    }

    public int getPreviousHash() {
        return previousHash;
    }

    public String[] getTransaction() {
        return transactions;
    }

    public int getBlockHash() {
        return blockHash;
    }
}
