import React, { useContext, useEffect, useState } from "react";
import { Deck } from "../types/deck";
import { useAuth } from "./auth";

import useAsyncEffect from 'use-async-effect';

export type Decks = {
    decks: Deck[],
    loading: boolean;
};

export type DecksCtx = Decks & {
    load: (deckId: number) => void;
};

const initialState: Decks = {
    decks: [],
    loading: false,
}

const DeckContext = React.createContext({
    ...initialState,
    load: (deckId: number) => {},
});

export function DecksProvider({children}: {children: React.ReactNode}) {
    const [state, setState] = useState<Decks>(initialState);
    const auth = useAuth();

    const load = async (deckId: number) => {
        let decks = state.decks;

        setState({...state, loading: true});
        
        for(let i = 0; i < decks.length; i++) {
            if (decks[i].id !== deckId) {
                continue;
            }

            if (!!decks[i].white_cards) {
                break;
            }

            const resp = await fetch(`/api/decks/${deckId}`);
            const data = await resp.json();

            decks[i].white_cards = data.white_cards;
            decks[i].black_cards = data.black_cards;

            break;
        }

        setState({...state, decks, loading: false});
    };

    useAsyncEffect(async () => {
        if (!auth.loaded) {
            return;
        }

        const resp = await fetch('/api/decks');
        const data = await resp.json();

        setState({...state, decks: data})
    }, [auth.loaded, auth.token]);

    return <DeckContext.Provider value={{
        ...state,
        load,
    }}>
        {children}
    </DeckContext.Provider>;
}

export function useDecks() {
    return useContext<DecksCtx>(DeckContext);
}