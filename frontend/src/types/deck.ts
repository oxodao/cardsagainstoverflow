export type Card = {
    id: number;
    text: string;
    is_black_card: boolean;
};

export type Deck = {
    id: number;
    name: string;
    author: string;
    selected_by_default: boolean;
    amt_white_cards: number;
    amt_black_cards: number;

    is_loaded: boolean;
    white_cards: Card[];
    black_cards: Card[];
};