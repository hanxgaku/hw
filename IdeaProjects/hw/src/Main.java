import java.util.Arrays;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        //TIP Press <shortcut actionId="ShowIntentionActions"/> with your caret at the highlighted text
        // to see how IntelliJ IDEA suggests fixing it.
        System.out.println("Hello and welcome!");
        int[] a = {1,3,5};
        int[] b = {2,4,6};
        int[] c = {2,1,3,5,4};
        System.out.println(Arrays.toString(join_arr(a, b)));
        }
    public static int factorial(int n) {
        int x = 1;
        while (n > 0) {
            x *= n;
            n--;
        }
        return x;
    }

    public static int factorial_rec(int n) {
        if (n <= 0) {
            return 1;
        }
        else {
            return factorial_rec(n - 1) * n;
        }
    }

    public static int factorial_rec2(int n, int x) {
        if (n > 0) {
            return factorial_rec2(n - 1, n * x);
        }
        else {
            return x;
        }
    }

    public static int pow(int n, int m) {
        int x = 1;
        for (int i = 0; i < m; i++) {
            x *= n;
        }
        return x;
    }

    public static int pow_rec(int n, int m) {
        if (m == 0) {
            return 1;
        }
        else {
            return n * pow_rec(n, m - 1);
        }
    }

    public static boolean is_sorted_array(int[] a) {
        for (int i = 0; i < a.length - 1; i++) {
            if (a[i] > a[i + 1]) {
                return false;
            }
        }
        return true;
    }

    public static boolean is_sorted_array_rec(int[] a, int index) {
        if ((index == a.length - 1) || (index == a.length)) {
            return true;
        }
        else if (a[index] > a[index + 1]) {
            return false;
        }
        else {
            return is_sorted_array_rec(a, index + 1);
        }
    }

    public static int[] reverse(int[] a) {
        int j, n;
        for (int i = 0; i < a.length / 2; i++) {
            j = a.length - i - 1;
            n = a[j];
            a[j] = a[i];
            a[i] = n;
        }
        return a;
    }

    public static int[] join_arr(int[] a, int[] b) {
        int i = 0, j = 0;
        int[] joined = new int[a.length + b.length];
        while ((i <= a.length) && (j <= b.length)) {
            if (i == a.length) {
                for (int k = j; k < b.length; k++) {
                    joined[i + k] = b[k];
                }
                break;
            }
            else if (j == b.length) {
                for (int k = i; k < a.length; k++) {
                    joined[i + k] = a[k];
                }
                break;
            }
            else if (a[i] < b[j]) {
                joined[i + j] = a[i];
                i++;
            }
            else {
                joined[i + j] = b[j];
                j++;
            }
        }
        return joined;
    }

    public static int[] join_arr_rec(int[] a, int[] b) {
        if (a.length == 0) {
            return b;
        }
        else if (b.length == 0) {
            return a;
        }
        else if (a[0] < b[0]) {
            int[] x = new int[a.length - 1];
            for (int i = 0; i < a.length - 1; i++) {
                x[i] = a[i + 1];
            }
            return join_arr(new int[] {a[0]}, join_arr_rec(x, b));
        }
        else {
            int[] x = new int[b.length - 1];
            for (int i = 0; i < b.length - 1; i++) {
                x[i] = b[i + 1];
            }
            return join_arr(new int[] {b[0]}, join_arr_rec(a, x));
        }
    }

    public static int[] bsort(int[] a) {
        int n, k = 0;
        for (int i = 0; i < a.length - 1; i++) {
            if (a[i] > a[i + 1]) {
                n = a[i];
                a[i] = a[i + 1];
                a[i + 1] = n;
                k++;
            }
        }
        if (k == 0) {
            return a;
        }
        else {
            return bsort(a);
        }
    }

    public static int[] bsort_while(int[] a) {
        while (true) {
            int n, k = 0;
            for (int i = 0; i < a.length - 1; i++) {
                if (a[i] > a[i + 1]) {
                    n = a[i];
                    a[i] = a[i + 1];
                    a[i + 1] = n;
                    k++;
                }
            }
            if (k == 0) {
                return a;
            }
        }
    }

    public static int[] bsort_while_1(int[] a) {
        while (true) {
            int n, k = 0, i = 0;
            while (i < a.length - 1) {
                if (a[i] > a[i + 1]) {
                    n = a[i];
                    a[i] = a[i + 1];
                    a[i + 1] = n;
                    k++;
                    if (i > 0) {
                        i = i - 2;
                    }
                }
                i++;
            }
            if (k == 0) {
                return a;
            }
        }
    }

    public static int[] msort(int[] a) {
        if (a.length == 0) {
            return new int[] {};
        }
        if (a.length == 1) {
            return new int[] {a[0]};
        }
        int k = a.length / 2;
        int[] x = new int[k], y = new int[a.length - k];
        for (int i = 0; i < k; i++) {
            x[i] = a[i];
        }
        for (int i = 0; i < a.length - k; i++) {
            y[i] = a[k + i];
        }
        System.out.println(Arrays.toString(x) + Arrays.toString(y));
        return join_arr(msort(x), msort(y));
    }

    public static int[] qsort(int[] a) {
        if (a.length == 0) {
            return new int[] {};
        }
        if (a.length == 1) {
            return new int[] {a[0]};
        }
        int k = a[0];
        int[] x = {}, y = {};
        for (int i = 1; i < a.length; i++) {
            if (a[i] < k) {
                x = join_arr(x, new int[] {a[i]});
            }
            else {
                y = join_arr(y, new int[] {a[i]});
            }
        }
        return join_arr(join_arr(qsort(x), new int[] {k}), qsort(y));
    }
}


