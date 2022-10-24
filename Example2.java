import java.io.*;

public class Example2 {

       public static void main(String[] args)

       throws IOException {

              if(args.length != 1) {

                     system.out.println("No arguments");

                     system.exit(1);

              }

              Runtime runtime = Runtime.getRuntime();

              String[] cmd = new String[3];

              cmd[0] = "cmd.exe" ;

                cmd[1] = "/C";

                cmd[2] = "dir " + args[0];

              Process proc = runtime.exec(cmd);

              

              InputStream is = proc.getInputStream();

              InputStreamReader isr = new InputStreamReader(is);

              BufferedReader br = new BufferedReader(isr);

              

              String line;

              while ((line = br.readLine()) != null) {

                     System.out.println(line);

              }

       }

}