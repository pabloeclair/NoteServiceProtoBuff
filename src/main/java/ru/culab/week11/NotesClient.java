package ru.culab.week11;


public class NotesClient implements NotesConnectable {

    @Override
    public void connectToServer(String hostName, int port) throws Exception {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'connectToServer'");
    }

    @Override
    public void close() {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'close'");
    }

    @Override
    public long createNote(String title, String content) throws Exception {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'createNote'");
    }

    @Override
    public String[] getNoteTitleAndContent(long id) throws Exception {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'getNoteTitleAndContent'");
    }

    @Override
    public void updateNote(long id, String title, String content) throws Exception {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'updateNote'");
    }

    @Override
    public void deleteNote(long id) throws Exception {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'deleteNote'");
    }

    @Override
    public long[] searchNotes(String pattern) throws Exception {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'searchNotes'");
    }


}