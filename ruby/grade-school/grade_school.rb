class School
  def initialize
    @grades = Hash.new { |hash, key| hash[key] = [] }
  end

  def add(name, grade)
    @grades[grade].insert(@grades[grade].bsearch_index { |n| n >= name } || -1, name)
  end

  def students(grade)
    @grades[grade]
  end

  def students_by_grade
    @grades.keys.sort.map { |grade| { grade: grade, students: students(grade) } }
  end
end

module BookKeeping
  VERSION = 3
end
