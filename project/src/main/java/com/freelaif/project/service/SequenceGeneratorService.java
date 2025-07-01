package com.freelaif.project.service;


import com.freelaif.project.models.Counter;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.FindAndModifyOptions;
import org.springframework.data.mongodb.core.MongoOperations;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.data.mongodb.core.query.Update;
import org.springframework.stereotype.Service;


@Service
public class SequenceGeneratorService {

    @Autowired
    private MongoOperations mongoOperations;

    public int getNextSequence(String seqName) {
        Query query = new Query(Criteria.where("_id").is(seqName));
        Update update = new Update().inc("seq", 1);
        FindAndModifyOptions options = new FindAndModifyOptions().returnNew(true).upsert(true);

        Counter counter = mongoOperations.findAndModify(query, update, options, Counter.class);
        return counter != null ? counter.getSeq() : 1;
    }
}
