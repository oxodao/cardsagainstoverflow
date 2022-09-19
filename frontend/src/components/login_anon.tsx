import {useTranslation} from 'react-i18next';
import '../assets/scss/login.scss';

export default function LoginAnonymously() {
    const {t} = useTranslation('translations', {keyPrefix: 'login'});

    return <form className="LoginForm">
        <div className="LoginForm__Display">
            <input id="LoginForm__RemoteDisplay" type="checkbox"/>
            <label htmlFor="LoginForm__RemoteDisplay">{t('remote_display')}</label>
        </div>
        
        <input type="text" id="username" name="username" placeholder={t('username')} required />
        <input type="text" id="room" name="room" placeholder={t('room_id')} />

        <input type="submit" value={t('create_join')} />
    </form>;
}