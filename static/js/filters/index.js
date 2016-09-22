import * as Constants from './constants.js';
import InvalidLinkError from './ErrorClass.js';
import Store from '../vuex/store';

/*
 * Parse and validate the url
 * 
 * This will check the url against some quite lenient rules
 * and try to only allow valid urls to pass but also allow people
 * to send through whatever crazy url they want to create.
 * 
 * All urls are returned URIEncoded
 * 
 * @param String input
 * @param Bool recursive
 * 
 * @returns String|Error 
 */
export function ValidateUrl(input, recursive = false){
        let parser = document.createElement('a');
        parser.href = input;
        
        const puncRegex = /[,.?\-]/;
        const relativeHostname = location.hostname;
        
        let href = parser.href;
        let proto = parser.protocol;
        let hostname = parser.hostname;
        let returnedURL;
        
        if(recursive && relativeHostname === hostname){
                throw new InvalidLinkError("Don't Shorten my urls");
        }
        
        /* Make sure there is a valid protocol and it's not too long' */
        if(proto.length > 0 && hostname.length > 0 && hostname.length < (126)
          && href.length < 2083){
                /* Make sure url is absolute and not relative */
                if(relativeHostname === hostname){
                        /* Add http and recursivly check if valid */
                        let absHostname = `${href.replace(relativeHostname + '/', '')}`;
                        return ValidateUrl(absHostname, true);
                }
                
                /* Hostnames cant start with a punctuation */
                if(puncRegex.test(hostname.replace('www.','').charAt(0))){
                        throw new InvalidLinkError("Links cant start with punctuation");
                }
                
                /* But they must have a period somewhere */
                if(hostname.indexOf('.') === -1){
                        throw new InvalidLinkError("The link contains an error");
                }
        }
        else{
                throw new InvalidLinkError("Link is of incorrect length");
        }
        
        returnedURL = href;
        parser = null;
        return returnedURL;
}

/*
 * Generate a key
 * 
 * @param int length
 * @return string
 */
function generateKey(length){
        let key = '';
        for(let i = 0; i < length; ++i){
                key += Constants.KEYSPACE[~~(Math.random() * (Constants.KEYSPACE.length))];
        }
        return key;
}

/*
 * Return a UID for a user
 * 
 * This will return a UID for a user either by finding the currently
 * set UID, or generating a new one
 * 
 * @retuns string|error
 */
export function GetUID(){
        if(window.localStorage.getItem(Constants.USERID_LOCALHOST) !== null){
                // Get already existing uid
                return window.localStorage.getItem(Constants.USERID_LOCALHOST);
        }
        else{
                // Generate new uid
                let key = generateKey(Constants.UID_LEN);
                // Set the key in localStorage
                window.localStorage.setItem(Constants.USERID_LOCALHOST, key);
                return key;
        }
}