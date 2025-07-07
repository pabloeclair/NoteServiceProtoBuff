package ru.culab.week11;

public interface NotesConnectable {
    
    // Connect to gRPC service
    void connectToServer(String hostName, int port) throws Exception;

    // Close gRPC connection and release resources
    void close();

    // Returns created note ID
    long createNote(String title, String content) throws Exception;

    // Returns 2-element array: 0:title and 1:content
    String[] getNoteTitleAndContent(long id) throws Exception;

    // Update existing note
    void updateNote(long id, String title, String content) throws Exception;

    // Delete note
    void deleteNote(long id) throws Exception;

    // Returns all matching note IDs
    long[] searchNotes(String pattern) throws Exception;
    
}
