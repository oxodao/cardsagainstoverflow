import {useTranslation} from 'react-i18next';
import '../assets/scss/login.scss';
import { useDecks } from '../hooks/decks';

export default function Login() {
    const {t} = useTranslation('translations', {keyPrefix: 'login'});

    const decks = useDecks();

    return <form className="LoginForm">
        <h1>{t('auth_title')}</h1>
        <input type="text" id="username" name="username" placeholder={t('username')} required />
        <input type="text" id="password" name="password" placeholder={t('password')} />

        <input type="submit" value={t('authenticate')} onClick={() => decks.load(2)} />
    </form>;
}