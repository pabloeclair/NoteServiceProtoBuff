package ru.culab.week11;

import java.util.List;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class NotesClient implements NotesConnectable {

    private ManagedChannel channel;
    private NoteServiceGrpc.NoteServiceBlockingStub stub;

    @Override
    public void connectToServer(String hostName, int port) throws Exception {
        ManagedChannel channel = ManagedChannelBuilder.forAddress(hostName, port).usePlaintext().build();
        this.channel = channel;
        this.stub = NoteServiceGrpc.newBlockingStub(channel);
    }

    @Override
    public void close() {
        this.channel.shutdown();
    }

    @Override
    public long createNote(String title, String content) throws Exception {
        NoteString req = NoteString.newBuilder().setName(title).setContent(content).build();
        NoteId resp = this.stub.createNote(req);
        return resp.getId();
    }

    @Override
    public String[] getNoteTitleAndContent(long id) throws Exception {
        NoteId req = NoteId.newBuilder().setId(id).build();
        NoteString resp = this.stub.getNote(req);
        String[] result = new String[2];
        result[0] = resp.getName();
        result[1] = resp.getContent();
        return result;
    }

    @Override
    public void updateNote(long id, String title, String content) throws Exception {
        UpdateNoteRequest req = UpdateNoteRequest.newBuilder()
            .setId(id)
            .setName(title)
            .setContent(content)
            .build();
        
        this.stub.updateNote(req);
    }

    @Override
    public void deleteNote(long id) throws Exception {
        NoteId req = NoteId.newBuilder().setId(id).build();
        this.stub.deleteNote(req);
    }

    @Override
    public long[] searchNotes(String pattern) throws Exception {
        SearchNotesRequest req = SearchNotesRequest.newBuilder().setPattern(pattern).build();
        NoteIdRepeated resp = this.stub.searchNotes(req);
        List<Long> result = resp.getIdList();
        return result.stream().mapToLong(l -> l).toArray();
    }


}