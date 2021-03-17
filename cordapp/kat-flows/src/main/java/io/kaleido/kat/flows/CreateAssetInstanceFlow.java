package io.kaleido.kat.flows;

import io.kaleido.kat.states.AssetInstanceCreated;
import net.corda.core.contracts.UniqueIdentifier;
import net.corda.core.flows.InitiatingFlow;
import net.corda.core.flows.StartableByRPC;
import net.corda.core.identity.Party;

import java.util.ArrayList;
import java.util.List;

@InitiatingFlow
@StartableByRPC
public class CreateAssetInstanceFlow extends CreateAssetEventFlow<AssetInstanceCreated>{
    private final String assetInstanceID;
    private final String assetDefinitionID;
    private final String contentHash;

    public CreateAssetInstanceFlow(String assetInstanceID, String assetDefinitionID, String contentHash, List<Party> observers) {
        super(observers);
        this.assetInstanceID = assetInstanceID;
        this.assetDefinitionID = assetDefinitionID;
        this.contentHash = contentHash;
    }

    @Override
    public AssetInstanceCreated getAssetEvent(){
        List<Party> participants = new ArrayList<>(this.observers);
        participants.add(getOurIdentity());
        return new AssetInstanceCreated(assetInstanceID, assetDefinitionID, getOurIdentity(), contentHash, participants);
    }
}