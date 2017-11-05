class School
  def initialize
    @grades = Hash.new { |hash, key| hash[key] = [] }
  end

  def add(name, grade)
    @grades[grade] << name
  end

  def students(grade)
    @grades[grade].sort
  end

  def students_by_grade
    @grades.keys.sort.map { |grade| { grade: grade, students: students(grade) } }
  end
end

module BookKeeping
  VERSION = 3
end
