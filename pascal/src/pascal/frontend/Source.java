package pascal.frontend;

import java.io.BufferedReader;
import java.io.IOException;

/**
 * <h1>Source</h1>
 *
 * <p>The framework class that represents the source program.</p>
 */

public class Source {
    public static final char EOL = '\n';            // end-of-line character
    public static final char EOF = (char) 0;        // end-of-file character

    private BufferedReader reader;                  // reader for the source program
    private String line;                            // source line
    private int lineNum;                            // current source line number
    private int currentPos;                         // current source line position

    /**
     * Constructor
     * @param reader the reader for the source program
     * @throws IOException if an I/O error occured
     */
    public Source(BufferedReader reader)
        throws IOException
    {
        this.lineNum = 0;
        this.currentPos = -2;
        this.reader = reader;
    }

    /**
     * Return the source character at the current position.
     * @return the source character at the current position
     * @throws Exception if an error occured
     */
    public char currentChar()
        throws Exception
    {
        // First time?
        if (currentPos == -2) {
            readLine();
            return nextChar();
        }

        // At the end of file?
        else if (line == null) {
            return EOF;
        }

        // At the end of line?
        else if ((currentPos == -1) || (currentPos == line.length())) {
            return EOL;
        }

        // Need to read the next line?
        else if (currentPos > line.length()) {
            readLine();
            return nextChar();
        }

        // Return the character at the current position.
        else {
            return line.charAt(currentPos);
        }
    }

    /**
     * Consume the current source character and return the next character
     * @return the next source character
     * @throws Exception if an error occured
     */
    public char nextChar()
        throws Exception
    {
        ++currentPos;
        return currentChar();
    }

    /**
     * Return the source character following the current character without
     * consuming the current character.
     * @return the foolowing character.
     * @throws Exception if an error occured.
     */
    public char peekChar()
        throws Exception
    {
        currentChar();
        if (line == null) {
            return EOF;
        }

        int nextPos = currentPos + 1;
        return nextPos < line.length() ? line.charAt(nextPos) : EOL;
    }

    /**
     * Read the next source line.
     * @throws IOException if an I/O error occurred.
     */
    private void readLine()
        throws IOException
    {
        line = reader.readLine();
        currentPos = -1;

        if (line != null) {
            ++lineNum;
        }
    }

    /**
     * Close the source.
     * @throws Exception if an error occured
     */
    public void close()
        throws Exception
    {
        if (reader != null) {
            try {
                reader.close();
            }
            catch (IOException ex) {
                ex.printStackTrace();
                throw ex;
            }
        }
    }
}
