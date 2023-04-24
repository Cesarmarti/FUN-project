from itertools import permutations
import math
import json 

def permutateSkill(skill):
    perms = list(permutations(skill))
    return set(perms)

def generateSkill(s_count,s_type,t_count,t_type):
    skill = []
    for i in range(s_count):
        skill.append("S"+s_type)
    for j in range(t_count):
        skill.append("T"+t_type)
    return skill 

def generateSkills(max_somersaults,max_twists):
    skills = []

    for i in range(1,max_somersaults+1):
        for svars in ["s","t","p"]:
            for j in range(1,max_twists+1):
                for tvars in ["s","t","p"]:
                    basic_skill = generateSkill(i,svars,j,tvars)
                    permutated_skills = list(permutateSkill(basic_skill))
                    skills = skills + permutated_skills

    return skills 


def calculateScore(skill):
    s_count = 0
    t_count = 0
    score = 0
    for ele in skill:
        if ele in ['Ss','Sp','St']:
            s_count = s_count + 1
        if ele in ['Ts','Tp','Tt']:
            t_count = t_count + 1
    score = score + 0.1*t_count
    score = score + 0.5*s_count
    s_count = s_count - 2
    if s_count > 0:
        score = score + s_count*0.1
    if t_count == 0:
        if s_count <= 2:
            score = score + 0.1*s_count

    score = math.floor(score * 100)/100.0
    return score

def skillsScore(skills):
    scores = []
    for skill in skills:
        score = calculateScore(skill)
        scores.append(score)
    return scores     

def print_skills(skills,scores):
    sport = {
        "discipline": "Skiing"
    }

    skills_array = []
    anti_rules = []
    for skill,score in zip(skills,scores):
        skill_name = ''.join(skill)
        skill_ele = {
            "name": skill_name,
            "label": skill_name,
            "value": score,
            "deduction": 0
        }
        anti_rule = {
            "k":1,
            "skills": [
                skill_name
            ]
        }
        anti_rules.append(anti_rule)
        skills_array.append(skill_ele)
    sport["skills"] = skills_array
    sport["antiRepetitionRule"] = {
        "groups": anti_rules
    }
    with open("trampoline.json", "w") as outfile:
        json.dump(sport, outfile)

if __name__ == '__main__':
    max_somersaults = 4 
    max_twists = 2
    skills = generateSkills(4,2)
    scores = skillsScore(skills)
    print_skills(skills,scores)