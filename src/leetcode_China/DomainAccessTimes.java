package leetcode_China;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * 输入:  ["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
 输出:  ["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
 说明:  按照假设，会访问"google.mail.com" 900次，"yahoo.com" 50次，"intel.mail.com" 1次，"wiki.org" 5次。
 而对于父域名，会访问"mail.com" 900+1 = 901次，"com" 900 + 50 + 1 = 951次，和 "org" 5 次。
 */

public class DomainAccessTimes {
    public List<String> subdomainVisits(String[] cpdomains) {
        List<String> list = new ArrayList<>();
        Map<String, Integer> timesMap = new HashMap<>();
        for(String s:cpdomains){
            String[] oneDomain = this.TimesAndDomain(s);
            toTimesMap(timesMap, oneDomain[0], oneDomain[1]);
        }
        for(Map.Entry entry:timesMap.entrySet()){
            list.add(entry.getValue()+" "+entry.getKey());
        }
        return list;
    }

    private String[] TimesAndDomain(String cpdomain){
        return cpdomain.split(" ");
    }

    private Map<String, Integer> toTimesMap(Map<String, Integer> timesMap, String time, String subDomain){
        Integer intTime = Integer.parseInt(time);
        String[] domains = subDomain.split("\\.");
        int n = domains.length;
        String midDomain = null;
        String topDomain = null;
        if(n == 3){
            midDomain = domains[1] + "." + domains[2];
            topDomain = domains[2];
            putDomainIntoMap(timesMap, midDomain, intTime);
        }else if (n == 2){
            topDomain = domains[1];
        }
        putDomainIntoMap(timesMap, subDomain, intTime);
        putDomainIntoMap(timesMap, topDomain, intTime);
        return timesMap;
    }

    private void putDomainIntoMap(Map<String, Integer> timesMap, String domain, int intTime){
        if(timesMap.containsKey(domain)){
            timesMap.put(domain, timesMap.get(domain) + intTime);
        }else {
            timesMap.put(domain, intTime);
        }
    }
}
