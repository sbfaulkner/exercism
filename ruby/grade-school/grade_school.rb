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
    @grades.sort.map { |grade, students| { grade: grade, students: students } }
  end
end

module BookKeeping
  VERSION = 3
end
